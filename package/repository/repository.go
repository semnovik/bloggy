package repository

import (
	"bloggy"
	"github.com/jmoiron/sqlx"
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

type Repository struct {
	Authorization
	Users
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Users:         NewUsersPostgres(db),
	}
}
