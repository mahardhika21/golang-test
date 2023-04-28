package repository

import "golang-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}