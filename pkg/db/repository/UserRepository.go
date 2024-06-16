package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"user-service/pkg/db/model"
)

type UserRepository struct {
	Connection *pgx.Conn
}

func mapUser(rows pgx.Rows, user *model.User) error {
	err := rows.Scan(&user.Id, &user.Name, &user.Surname, &user.Email, &user.PasswordHash, &user.GroupId)
	if err != nil {
		return err
	}
	return nil
}

const QUERY_GET_USER_BY_ID = "SELECT id, name, surname, email, password_hash, group_id FROM app_users WHERE id=$1"

func (userRepository *UserRepository) GetUserById(id int64) (error, *model.User) {
	if userRepository.Connection == nil {
		return errors.New("connection is nil"), nil
	}
	rows, err := userRepository.Connection.Query(context.Background(), QUERY_GET_USER_BY_ID, id)
	if err != nil {
		return err, nil
	}
	user := new(model.User)
	err = mapUser(rows, user)
	if err != nil {
		return err, nil
	}
	return nil, user
}
