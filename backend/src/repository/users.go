package repository

import "github.com/traP-jp/h23s_01/backend/src/domain"

type UsersRepository interface {
	RemakeUsersTable(users []domain.User) error
}
