package usecase

import (
	"errors"
	"net/mail"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/rizkyimaduddin24/techtest/internal/entity"
	"gitlab.com/rizkyimaduddin24/techtest/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(name string, email string, password string) error
	Login(email string, password string) (uint, string, error)
}

type authUsecase struct {
	repo repository.UserRepository
}

func NewAuthUsecase(repo repository.UserRepository) AuthUsecase {
	return &authUsecase{
		repo: repo,
	}
}

func (uc *authUsecase) Register(name string, email string, password string) error {
	// Validate Email
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email")
	}

	user := &entity.User{}
	err = uc.repo.GetByEmail(user, email)
	if user.ID != 0 || err != nil {
		return errors.New("email already exist")
	}

	// Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Create User
	user = &entity.User{
		Name:     name,
		Email:    email,
		Role:     "user",
		Password: string(hash),
	}
	if err := uc.repo.Create(user); user.ID == 0 || err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (uc *authUsecase) Login(email string, password string) (uint, string, error) {
	// Get User
	user := &entity.User{}
	err := uc.repo.GetByEmail(user, email)
	if user.ID == 0 || err != nil {
		return 0, "", errors.New("user not found")
	}

	// Compare Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return 0, "", errors.New("invalid password")
	}

	// Make Claims
	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign Token
	secret := os.Getenv("JWT_SECRET")
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return 0, "", errors.New("failed to sign token")
	}
	signed = "Bearer " + signed

	return user.ID, signed, nil
}
