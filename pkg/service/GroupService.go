package service

import (
	"user-service/pkg/db"
	"user-service/pkg/db/model"
)

type GroupService struct {
	RepositoryFactoryBean *db.RepositoryFactory
}

func (service *GroupService) InitService(factory *db.RepositoryFactory) {
	service.RepositoryFactoryBean = factory
}

func (service *GroupService) SearchGroupById(id int64) (error, *model.Group) {
	err, group := service.RepositoryFactoryBean.GroupRepositoryBean.GetGroupById(id)
	if err != nil {
		return err, nil
	}
	return nil, group
}
