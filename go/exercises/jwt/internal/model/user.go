package model

import "fmt"

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	IsActive  bool   `json:"is_active"`
}

func (u User) Representation() {
	fmt.Printf(
		"User(%d, %s, %s, %v, %d, %v)",
		u.ID, u.Name, u.Email, u.CreatedAt, u.UpdatedAt, u.IsActive,
	)
}
