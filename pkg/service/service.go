package service

import (
	serverapp "RestApp"
	"RestApp/pkg/repository"
)

type Authorization interface {
	CreateUser(user serverapp.User) (int, error)
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

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
