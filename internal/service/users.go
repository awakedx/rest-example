package service

import (
	"context"
	"fmt"
	"os"
	"time"
	"webproj/internal/domain"
	"webproj/internal/repository"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	repo           repository.Users
	accessTokenTTL time.Duration
}

func NewUsersService(repo repository.Users, accessTokenTTL time.Duration) *UsersService {
	return &UsersService{
		repo:           repo,
		accessTokenTTL: accessTokenTTL,
	}
}

func (s *UsersService) GetAccTokenTTL() time.Duration {
	return s.accessTokenTTL
}

func (s *UsersService) SignUp(ctx context.Context, input *SignUpInput) error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return err
	}
	user := domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  string(hashPass),
		CreatedAt: time.Now(),
	}
	if err := s.repo.Create(ctx, &user); err != nil {
		return err
	}
	return nil
}

func (s *UsersService) SignIn(ctx context.Context, input *SignInInput) (string, error) {
	user, err := s.repo.Get(ctx, input.Email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", fmt.Errorf("pass is wrong")
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": user.Id,
			"exp":    time.Now().Add(s.accessTokenTTL).Unix(),
		})
	token, err := tokenClaims.SignedString([]byte(os.Getenv("SECRET_ACCESS")))
	if err != nil {
		return "", err
	}
	return token, err
}

func (s *UsersService) DeleteUser(ctx context.Context, userId uuid.UUID) error {
	if err := s.repo.Delete(ctx, userId); err != nil {
		return err
	}
	return nil
}

func (s *UsersService) GetById(ctx context.Context, userId uuid.UUID) (*domain.User, error) {
	user, err := s.repo.Get(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
