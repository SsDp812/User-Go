package db

import (
	"errors"
	"user-service/pkg/db/repository"
)

type RepositoryFactory struct {
	connectionFactory   *ConnectionFactory
	UserRepositoryBean  *repository.UserRepository
	GroupRepositoryBean *repository.GroupRepository
}

func InitRepositoryFactory(connectionFactoryInit *ConnectionFactory) (error, *RepositoryFactory) {
	if connectionFactoryInit == nil {
		return errors.New("connectionFactory is nil"), nil
	}
	repositoryFactory := &RepositoryFactory{connectionFactory: connectionFactoryInit}
	repositoryFactory.UserRepositoryBean = &repository.UserRepository{Connection: connectionFactoryInit.Connect}
	repositoryFactory.GroupRepositoryBean = &repository.GroupRepository{Connection: connectionFactoryInit.Connect}
	return nil, repositoryFactory
}
