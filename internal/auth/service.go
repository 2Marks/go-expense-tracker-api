package auth

import (
	"fmt"

	"github.com/2marks/go-expense-tracker-api/types"
	"github.com/2marks/go-expense-tracker-api/utils"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo types.AuthRepository
}

func NewService(repo types.AuthRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Signup(params types.SignupDTO) error {
	if isUsernameExist := s.repo.IsUsernameExist(params.Username); isUsernameExist {
		return fmt.Errorf("user with username:%s already exist", params.Username)
	}

	if isEmailExist := s.repo.IsEmailExist(params.Email); isEmailExist {
		return fmt.Errorf("user with email:%s already exist", params.Email)
	}

	passwordHash := s.generatePasswordHash(params.Password)

	if passwordHash == "" {
		return fmt.Errorf("error occured while creating user account. please contact admin")
	}

	return s.repo.CreateUser(&types.User{
		Name:     params.Name,
		Username: params.Username,
		Email:    params.Email,
		Password: passwordHash,
		IsActive: true,
	})
}

func (s *Service) Login(params types.LoginDTO) (*types.LoginResponse, error) {
	user, err := s.repo.GetUserByUsername(params.Username)

	if err != nil {
		return nil, fmt.Errorf("invalid login credentials")
	}

	isPasswordValid := s.validatePassword(user.Password, params.Password)
	if !isPasswordValid {
		return nil, fmt.Errorf("invalid login credentials")
	}

	authToken, err := utils.GenerateAuthToken(int(user.ID))
	if err != nil {
		fmt.Printf("error while generating auth token. err:%s", err.Error())
		return nil, fmt.Errorf("error occurred while loggin in. please contact admin")
	}

	loginResponse := types.LoginResponse{
		ID:    int64(user.ID),
		Token: authToken,
	}

	return &loginResponse, nil
}

func (s *Service) generatePasswordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Printf("error while generating password hash: %s", err.Error())
		return ""
	}

	return string(hash)
}

func (s *Service) validatePassword(hashed string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))

	return err == nil
}
