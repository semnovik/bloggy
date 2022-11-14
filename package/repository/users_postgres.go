package repository

import (
	"bloggy"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UsersPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{db: db}
}

func (r *UsersPostgres) UsersList() ([]bloggy.SingleUser, error) {
	var users []bloggy.SingleUser
	query := fmt.Sprintf("SELECT id, name, username from %s", usersTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each bloggy.SingleUser
		err = rows.Scan(&each.Id, &each.Name, &each.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, each)
	}

	return users, err
}
