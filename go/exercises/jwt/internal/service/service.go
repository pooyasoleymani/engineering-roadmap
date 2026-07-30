package service

import "gostudy/internal/model"

type Service interface {
	Login(username, password string) error
	Register(username, password string) error
	GetProfile(tokenString string) (model.Model, error)
}
