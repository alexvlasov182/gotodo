package service

import (
	"github.com/alexvlasov182/gotodo/pkg/repository"
	"github.com/alexvlasov182/gotodo"
	"crypto/sha1"
	"fmt"
)

const salt = "hasdflkasdflk123123kj1k23"

type AuthService struct{
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
