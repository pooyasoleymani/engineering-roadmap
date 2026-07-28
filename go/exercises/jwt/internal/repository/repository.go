package repository

import "gostudy/internal/model"

type Repository interface {
	GetAll() ([]model.Model, error)
	GetByID(id int64) (model.Model, error)
	Update(id int64) error
	Delete(id int64) error
}
