package services

import (
	"errors"
	"go-unit-test/entity"
	"go-unit-test/repository"
)

type CategoryServices struct {
	Repository repository.CategoryRepository
}

func (services CategoryServices) Get(id string) (*entity.Category, error)  {
	category := services.Repository.FindById(id)

	if category == nil {
		return nil, errors.New("Category not found")
	}else {
		return category, nil
	}
}