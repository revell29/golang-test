package repository

import "dira-go/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}