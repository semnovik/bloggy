package service

import (
	"bloggy"
	"bloggy/package/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) UsersList() ([]bloggy.SingleUser, error) {
	return s.repo.UsersList()
}
