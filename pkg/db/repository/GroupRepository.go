package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"user-service/pkg/db/model"
)

type GroupRepository struct {
	Connection *pgx.Conn
}

func mapGroup(rows pgx.Rows, group *model.Group) error {
	err := rows.Scan(&group.GroupId, &group.GroupName)
	if err != nil {
		return err
	}
	return nil
}

const QUERY_GET_GROUP_BY_ID = "SELECT id, name FROM app_user_groups WHERE id=$1"

func (groupRepository *GroupRepository) GetGroupById(id int64) (error, *model.Group) {
	if groupRepository.Connection == nil {
		return errors.New("connection is nil"), nil
	}
	rows, err := groupRepository.Connection.Query(context.Background(), QUERY_GET_GROUP_BY_ID, id)
	if err != nil {
		return err, nil
	}
	group := new(model.Group)
	err = mapGroup(rows, group)
	if err != nil {
		return err, nil
	}
	return nil, group
}
