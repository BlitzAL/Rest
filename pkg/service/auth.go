package service

import (
	serverapp "RestApp"
	"RestApp/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const salt = "dsaasd"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repo: repos}
}

func (s *AuthService) CreateUser(user serverapp.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
