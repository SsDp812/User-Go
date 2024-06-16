package service

import (
	"user-service/pkg/db"
	"user-service/pkg/db/model"
)

type UserService struct {
	RepositoryFactoryBean *db.RepositoryFactory
}

func (service *UserService) InitService(factory *db.RepositoryFactory) {
	service.RepositoryFactoryBean = factory
}

func (service *UserService) SearchUserById(id int64) (error, *model.User) {
	err, user := service.RepositoryFactoryBean.UserRepositoryBean.GetUserById(id)
	if err != nil {
		return err, nil
	}
	return nil, user
}
