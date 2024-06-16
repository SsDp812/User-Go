package service

import (
	"errors"
	"user-service/pkg/db"
)

type Service interface {
	InitService(factory *db.RepositoryFactory)
}

func InitializeAllServices(factory *db.RepositoryFactory, services ...Service) error {
	if factory == nil {
		return errors.New("factory cannot be nil")
	}
	for _, service := range services {
		service.InitService(factory)
	}
	return nil
}
