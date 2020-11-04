package repository

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/line_notify/src/db"
	"github.com/line_notify/src/model/user"
)

type UsersRepository interface {
	GetUsers(ctx context.Context) ([]user.User, error)
}

func NewUsersRepository(db *db.DBContext) UsersRepository {
	return usersRepositoryImpl{
		db: db.DB,
	}
}

type usersRepositoryImpl struct {
	db *sql.DB
}

func (u usersRepositoryImpl) GetUsers(ctx context.Context) ([]user.User, error) {
	rows, err := u.db.Query("SELECT id, name FROM m_users ORDER BY id;")
	if err != nil {
		return nil, err
	}
	var list []user.User
	for rows.Next() {
		var user user.User
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		list = append(list, user)
	}
	return list, nil
}
