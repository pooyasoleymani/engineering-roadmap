package repository

import (
	"crypto/sha3"
	"fmt"
	"gostudy/internal/model"
	"log"
	"slices"
	"sync/atomic"
	"time"
)

type UserRepository struct {
	Db     []*model.User
	logger *log.Logger
	id     *atomic.Int64
}

func NewUserRepository(db []*model.User, logger *log.Logger) *UserRepository {
	id := &atomic.Int64{}
	id.Store(int64(0))
	return &UserRepository{
		db,
		logger,
		id,
	}
}

func (r *UserRepository) GetAll() ([]*model.User, error) {

	return r.Db, nil
}

func (r *UserRepository) GetByID(id int64) (*model.User, error) {
	for _, user := range r.Db {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (r *UserRepository) Update(id int64, name, email, password string) error {
	for _, user := range r.Db {
		if user.ID == id {
			user.Name = name
			if name != "" {
				user.Name = name
			}
			if email != "" {
				user.Email = email
			}
			if password != "" {
				hash := sha3.Sum256([]byte(password))
				user.Password = string(hash[:])

			}
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

func (r *UserRepository) Delete(id int64) error {
	index := -1
	for _, user := range r.Db {
		index++
		if user.ID == id {
			slices.Delete(r.Db, index, index)
		}
	}
	return fmt.Errorf("cant delete user %d", id)
}

func (r *UserRepository) Create(username, password, email string) error {
	hash := sha3.Sum256([]byte(password))

	r.Db = append(r.Db, &model.User{
		ID:        r.id.Add(1),
		Name:      username,
		Email:     email,
		Password:  string(hash[:]),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})
	return nil
}

func (r *UserRepository) GetByUserAndPassword(username, password string) (*model.User, error) {
	hash := sha3.Sum256([]byte(password))
	users, err := r.GetAll()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Name == username && user.Password == string(hash[:]) {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}
