package service

import (
	"bloggy"
	"bloggy/package/repository"
)

type Authorization interface {
	CreateUser(user bloggy.User) (int, string, error)
}

type Users interface {
	UsersList() ([]bloggy.SingleUser, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	Users
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Users:         NewUsersService(repos.Users),
	}
}
