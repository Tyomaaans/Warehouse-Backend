package redis

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"

	"Backend-Warehouse/domain/entity"
	"Backend-Warehouse/domain/repository"
)

// ─── Konstanta ────────────────────────────────────────────────────────────────

const (
	refreshTokenPrefix = "refresh:"
	sessionPrefix      = "session:"
	refreshTokenLength = 32
)

// ─── Internal Types ───────────────────────────────────────────────────────────

// jwtClaims adalah struktur internal untuk signing JWT.
// Tidak di-ekspos keluar; domain menggunakan entity.Claims.
type jwtClaims struct {
	EmployeeID string `json:"employee_id"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	gojwt.RegisteredClaims
}

// refreshTokenPayload adalah data yang disimpan di Redis per refresh token.
type refreshTokenPayload struct {
	EmployeeID string    `json:"employee_id"`
	Username   string    `json:"username"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiresAt  time.Time `json:"expires_at"`
}

// ─── Service ──────────────────────────────────────────────────────────────────

type jwtService struct {
	secretKey          string
	accessTokenExpiry  time.Duration
	refreshTokenExpiry time.Duration
	redisClient        *redis.Client
}

// NewJWTService membuat instance jwtService yang mengimplementasikan repository.JWTService.
func NewJWTService(
	secretKey string,
	accessTokenExpiry time.Duration,
	refreshTokenExpiry time.Duration,
	redisClient *redis.Client,
) repository.JWTService {
	return &jwtService{
		secretKey:          secretKey,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
		redisClient:        redisClient,
	}
}

// ─── Access Token ─────────────────────────────────────────────────────────────

func (s *jwtService) generateAccessToken(employeeID, username, role string) (string, error) {
	now := time.Now()
	claims := &jwtClaims{
		EmployeeID: employeeID,
		Username:   username,
		Role:       role,
		RegisteredClaims: gojwt.RegisteredClaims{
			ExpiresAt: gojwt.NewNumericDate(now.Add(s.accessTokenExpiry)),
			IssuedAt:  gojwt.NewNumericDate(now),
		},
	}
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

// parseAccessToken mem-parse token tanpa validasi expiry.
// Validasi expiry dilakukan secara manual di ValidateAccessToken
// agar urutan pengecekan bisa dikontrol (session check dulu, baru expiry).
func (s *jwtService) parseAccessToken(tokenStr string) (*jwtClaims, error) {
	claims := &jwtClaims{}
	_, err := gojwt.ParseWithClaims(
		tokenStr,
		claims,
		func(t *gojwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*gojwt.SigningMethodHMAC); !ok {
				return nil, repository.ErrTokenInvalid
			}
			return []byte(s.secretKey), nil
		},
		gojwt.WithoutClaimsValidation(),
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", repository.ErrTokenInvalid, err)
	}
	return claims, nil
}

// ─── Session ──────────────────────────────────────────────────────────────────
//
// Pendekatan session-based menghindari gap TTL pada blacklist per-token.
// Session dibuat saat login (TTL = refreshTokenExpiry) dan dihapus saat logout.
// Dengan begitu token yang sudah expired pun tetap di-revoke setelah logout.

func (s *jwtService) createSession(ctx context.Context, userId string) error {
	return s.redisClient.Set(ctx, sessionPrefix+userId, "active", s.refreshTokenExpiry).Err()
}

func (s *jwtService) deleteSession(ctx context.Context, userId string) error {
	return s.redisClient.Del(ctx, sessionPrefix+userId).Err()
}

func (s *jwtService) isSessionActive(ctx context.Context, userId string) (bool, error) {
	val, err := s.redisClient.Get(ctx, sessionPrefix+userId).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return val == "active", nil
}

// ─── Refresh Token ────────────────────────────────────────────────────────────

func (s *jwtService) generateRefreshToken(ctx context.Context, employeeID, username, role string) (string, error) {
	b := make([]byte, refreshTokenLength)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to generate refresh token: %w", err)
	}
	token := hex.EncodeToString(b)

	now := time.Now()
	payload := refreshTokenPayload{
		EmployeeID: employeeID,
		Username:   username,
		Role:       role,
		CreatedAt:  now,
		ExpiresAt:  now.Add(s.refreshTokenExpiry),
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal refresh token payload: %w", err)
	}

	key := refreshTokenPrefix + token
	if err := s.redisClient.Set(ctx, key, data, s.refreshTokenExpiry).Err(); err != nil {
		return "", fmt.Errorf("failed to store refresh token: %w", err)
	}

	return token, nil
}

func (s *jwtService) getRefreshTokenPayload(ctx context.Context, refreshToken string) (*refreshTokenPayload, error) {
	data, err := s.redisClient.Get(ctx, refreshTokenPrefix+refreshToken).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, repository.ErrRefreshTokenInvalid
	}
	if err != nil {
		return nil, fmt.Errorf("redis error: %w", err)
	}

	var payload refreshTokenPayload
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, fmt.Errorf("corrupted refresh token data: %w", err)
	}
	return &payload, nil
}

// ─── Public Methods (implementasi repository.JWTService) ─────────────────────

// GenerateTokenPair membuat access + refresh token baru, lalu memperbarui session.
func (s *jwtService) GenerateTokenPair(ctx context.Context, userId, username, role string) (*entity.TokenPair, error) {
	accessToken, err := s.generateAccessToken(userId, username, role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := s.generateRefreshToken(ctx, userId, username, role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	if err := s.createSession(ctx, userId); err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return &entity.TokenPair{
		EmployeeID:   userId,
		Username:     username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// ValidateAccessToken memvalidasi token dengan urutan:
//  1. Parse & verifikasi signature
//  2. Cek session aktif (jika tidak ada → sudah logout)
//  3. Cek expiry
func (s *jwtService) ValidateAccessToken(ctx context.Context, tokenStr string) (*entity.Claims, error) {
	claims, err := s.parseAccessToken(tokenStr)
	if err != nil {
		return nil, repository.ErrTokenInvalid
	}

	active, err := s.isSessionActive(ctx, claims.EmployeeID)
	if err != nil {
		return nil, repository.ErrTokenBlocked
	}
	if !active {
		return nil, repository.ErrTokenRevoked
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		return nil, repository.ErrTokenExpired
	}

	return &entity.Claims{
		EmployeeID: claims.EmployeeID,
		Username:   claims.Username,
		Role:       claims.Role,
		ExpiresAt:  claims.ExpiresAt.Time,
		IssuedAt:   claims.IssuedAt.Time,
	}, nil
}

// RefreshTokens menukar refresh token lama dengan token pair baru.
// Refresh token lama langsung dihapus (token rotation) untuk mencegah reuse.
func (s *jwtService) RefreshTokens(ctx context.Context, refreshToken string) (*entity.TokenPair, error) {
	payload, err := s.getRefreshTokenPayload(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	if err := s.redisClient.Del(ctx, refreshTokenPrefix+refreshToken).Err(); err != nil {
		return nil, fmt.Errorf("failed to invalidate old refresh token: %w", err)
	}

	return s.GenerateTokenPair(ctx, payload.EmployeeID, payload.Username, payload.Role)
}

// RevokeTokens menghapus session dan refresh token sekaligus saat logout.
// Access token tidak perlu di-blacklist karena session-based check sudah cukup.
func (s *jwtService) RevokeTokens(ctx context.Context, accessToken, refreshToken string) error {
	claims, err := s.parseAccessToken(accessToken)
	if err != nil {
		return repository.ErrTokenInvalid
	}

	if err := s.deleteSession(ctx, claims.EmployeeID); err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	if err := s.redisClient.Del(ctx, refreshTokenPrefix+refreshToken).Err(); err != nil {
		return fmt.Errorf("failed to revoke refresh token: %w", err)
	}

	return nil
}