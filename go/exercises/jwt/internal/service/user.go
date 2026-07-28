package service

import (
	"errors"
	"fmt"
	"gostudy/internal/model"
	"gostudy/internal/repository"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	Repository *repository.UserRepository
	logger     *log.Logger
}

func NewUserService(repo *repository.UserRepository, logger *log.Logger) *UserService {
	return &UserService{
		Repository: repo,
		logger:     logger,
	}
}

func (u *UserService) Register(username, password string) error {
	if username == "" {
		return errors.New("username is empty")
	}
	if password == "" {
		return errors.New("password is empty")
	}
	users, err := u.Repository.GetAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		if username == user.Name {
			return errors.New("username already exists")
		}
	}

	return u.Repository.Create(username, password, "")
}

func (u *UserService) GetProfile(tokenString string) (*model.User, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	id := token.Header["user_id"].(int64)
	users, err := u.Repository.GetAll()
	if err != nil {
		return nil, err
	}
	if token.Header["exp"].(int64) < time.Now().Unix() {
		return nil, fmt.Errorf("token expired")
	}

	for _, user := range users {
		if id == user.ID {
			u.logger.Println("token valid")
			return u.Repository.GetByID(id)
		}
	}
	return nil, fmt.Errorf("user not exists")
}

func (u *UserService) Login(username, password string) (*model.User, error) {
	return u.Repository.GetByUserAndPassword(username, password)
}
