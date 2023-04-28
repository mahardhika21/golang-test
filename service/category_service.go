package service

import (
	"golang-test/repository"
	"golang-test/entity"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if Category == nil {
		retur category, errors.New("Category not found")
	} else {
		return category, nil
	}
}