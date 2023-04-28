package service

import (
		"testing"
		"golang-test/repository"
		"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRpositoryMock{Mock : mock.Mock {}}
var categoryService = CategoryService{Repository: CategoryRepository}

func TestCategoryService_Get(t *testing.T) {
	// program moc
	categoryRepository.Mock.On("FindBy", "1").Return(nil)

	category, err = categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
} 