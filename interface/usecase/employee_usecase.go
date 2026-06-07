package usecase

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	"Backend-Warehouse/domain/entity"
	"Backend-Warehouse/domain/repository"
	"Backend-Warehouse/interface/dto"
	"Backend-Warehouse/interface/mapper"
	jsonValidator "Backend-Warehouse/validator"
)

type EmployeeUseCase struct {
	empRepo    repository.EmployeeRepository
	jwtService repository.JWTService
	validate   *validator.Validate
}

func NewEmployeeUseCase(
	empRepo    repository.EmployeeRepository,
	jwtService repository.JWTService,
	validate   *validator.Validate,
) *EmployeeUseCase {
	return &EmployeeUseCase{
		empRepo:    empRepo,
		jwtService: jwtService,
		validate:   validate,
	}
}

func (uc *EmployeeUseCase) RegisterEmployee(ctx context.Context, req dto.RegisterEmployeeRequest) error {
	if err := jsonValidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	if req.Password != req.ConfirmPassword {
		return errors.New("password do not match!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed hash password!")
	}

	req.Password = string(hashedPassword)

	emp := mapper.ToEntityFromRegisterEmployeeRequest(req)
	if err := uc.empRepo.CreateEmployee(ctx, emp); err != nil {
		return err
	}

	return nil
}

func (uc *EmployeeUseCase) LoginEmployee(ctx context.Context, req dto.LoginEmployeeRequest) (*dto.LoginEmployeeResponse, *entity.TokenPair, error) {
	if err := jsonValidator.ValidateStruct(uc.validate, req); err != nil {
		return nil, nil, err
	}

	var emp *entity.EmployeeEntity
	var err error
	var ErrInvalidCredentials = errors.New("invalid credentials!")

	if req.Username != "" {
		emp, err = uc.empRepo.GetEmployeeByEmployeeUsername(ctx, req.Username)
	}
	if req.Email != "" {
		emp, err = uc.empRepo.GetEmployeeByEmployeeEmail(ctx, req.Email)
	}
	if req.Username == "" && req.Email == "" {
		return nil, nil, errors.New("username or email required!")
	}

	if err != nil {
		if errors.Is(err, repository.HandleDBError(err)) {
			return nil, nil, ErrInvalidCredentials
		}
		return nil, nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(emp.Password), []byte(req.Password)); err != nil {
		return nil, nil, ErrInvalidCredentials
	}

	token, err := uc.jwtService.GenerateTokenPair(ctx, emp.EmployeeID, emp.Username, emp.Role)
	if err != nil {
		return nil, nil, err
	}

	res := dto.LoginEmployeeResponse{
		AccessToken: token.AccessToken,
		EmployeProfile: mapper.ToEmployeeResponse(emp),
	}

	return &res, token, nil
}

func (uc *EmployeeUseCase) RefreshToken(ctx context.Context, refreshToken string) (*dto.RefreshTokenResponse, error) {
	pairToken, err := uc.jwtService.RefreshTokens(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	return &dto.RefreshTokenResponse{
		AccessToken:  pairToken.AccessToken,
		RefreshToken: pairToken.RefreshToken,
	}, nil
}

func (uc *EmployeeUseCase) LogoutEmployee(ctx context.Context, accessToken, refreshToken string) error {
	if err := uc.jwtService.RevokeTokens(ctx, accessToken, refreshToken); err != nil {
		return err
	}

	return nil
}

func (uc *EmployeeUseCase) GetEmployee(ctx context.Context, employeeID string) (*dto.EmployeeResponse, error) {
	emp, err := uc.empRepo.GetEmployeeByEmployeeID(ctx, employeeID)
	if err != nil {
		return nil, err
	}

	res := mapper.ToEmployeeResponse(emp)

	return &res, nil
}

func (uc *EmployeeUseCase) GetEmployees(ctx context.Context) ([]dto.EmployeeResponse, error) {
	emps, err := uc.empRepo.GetEmployees(ctx)
	if err != nil {
		return nil, err
	}

	res := mapper.ToEmployeeResponseList(emps)

	return res, nil
}

func (uc *EmployeeUseCase) UpdateEmployeeForUser(ctx context.Context, req dto.UpdateEmployeeForUserRequest) error {
	if err := jsonValidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	update := mapper.ToUpdateEmployeeEntityForUser(req)
	if err := uc.empRepo.UpdateEmployeeForUser(ctx, update); err != nil {
		return err
	}

	return nil
}

func (uc *EmployeeUseCase) UpdateEmployeeForAdmin(ctx context.Context, req dto.UpdateEmployeeForAdminRequest) error {
	if err := jsonValidator.ValidateStruct(uc.validate, req); err != nil {
		return  err
	}

	if req.Password != nil && *req.Password != "" {
		if req.Password != req.ConfirmPassword {
			return errors.New("password do not match!")
		}
	}

	update := mapper.ToUpdateEmployeeEntityForAdmin(req)
	if err := uc.empRepo.UpdateEmployeeForAdmin(ctx, update); err != nil {
		return err
	}

	return nil
}