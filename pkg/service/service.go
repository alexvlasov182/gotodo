package service

import (
	"github.com/alexvlasov182/gotodo/pkg/repository"
	"github.com/alexvlasov182/gotodo"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {

}

type TodoItem interface {

}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
