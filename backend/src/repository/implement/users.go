package implement

import (
	"github.com/jmoiron/sqlx"
	"github.com/traP-jp/h23s_01/backend/src/domain"
)

type Users struct {
	db *sqlx.DB
}

func NewUsers(db *sqlx.DB) *Users {
	return &Users{
		db: db,
	}
}

func (u *Users) RemakeUsersTable(users []domain.User) error {
	u.db.MustExec("DELETE FROM users")
	_, err := u.db.NamedExec("INSERT INTO users (id, name) VALUES (:id, :name)", users)
	if err != nil {
		return err
	}
	return nil
}

func (u *Users) GetUserNameById(id string) (string, error) {
	var userName string
	err := u.db.Get(&userName, "SELECRT name FROM users WHERE id = ?", id)
	if err != nil {
		return "", err
	}
	return userName, nil
}
