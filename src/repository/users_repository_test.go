package repository

import (
	"context"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/line_notify/src/model/user"
)

func TestUsersRepository_GetUsers(t *testing.T) {
	db := InitDB()
	ur := NewUsersRepository(db)
	ctx := context.TODO()

	tests := []struct {
		name string
		err  error
		want []user.User
	}{
		{
			name: "getUsers",
			err:  nil,
			want: []user.User{
				{
					ID:   1,
					Name: "test user",
				},
				{
					ID:   2,
					Name: "test user2",
				},
				{
					ID:   3,
					Name: "test user3",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := ur.GetUsers(ctx)

			if test.err != nil {
				return
			}

			assert.Equal(t, test.want, result)
		})
	}
	CloseDB(db)
}
