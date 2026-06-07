package repository

import (
	"context"
	"errors"

	"Backend-Warehouse/domain/entity"
)

var (
	ErrTokenExpired         = errors.New("access token expired")
	ErrTokenRevoked         = errors.New("token has been revoked")
	ErrTokenInvalid         = errors.New("invalid token")
	ErrTokenBlocked         = errors.New("failed to check token status")
	ErrRefreshTokenInvalid  = errors.New("refresh token expired or already used, please login again")
	ErrRefreshTokenNotFound = errors.New("refresh token not found")
)

type JWTService interface {
	GenerateTokenPair(ctx context.Context, employeeID, username, role string) (*entity.TokenPair, error)
	ValidateAccessToken(ctx context.Context, tokenStr string) (*entity.Claims, error)
	RefreshTokens(ctx context.Context, refreshToken string) (*entity.TokenPair, error)
	RevokeTokens(ctx context.Context, accessToken, refreshToken string) error
}