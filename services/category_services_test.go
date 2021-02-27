package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-unit-test/entity"
	"go-unit-test/repository"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryServices = CategoryServices{Repository: categoryRepository}

func TestCategoryServices_GetNotFound(t *testing.T) {
	//program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryServices.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryServices_GetSucces(t *testing.T) {
	category := entity.Category{
		Id: "1",
		Name: "Laptop",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryServices.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
