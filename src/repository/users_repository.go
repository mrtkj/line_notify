package repository

import (
	"context"
	"os/user"

	"github.com/line_notify/src/db"
)

type UsersRepository interface {
	GetUsers(ctx context.Context) ([]user.User, error)
}

func NewUsersRepository(db *db.DBContext) UsersRepository {
	return usersRepositoryImpl{
		db: db,
	}
}

type usersRepositoryImpl struct {
	db *db.DBContext
}

func (u usersRepositoryImpl) GetUsers(ctx context.Context) ([]user.User, error) {
	return nil, nil
}
