package http_handler

import (
	"gostudy/internal/service"
	"log"
)

type UserHTTPHandler struct {
	userService *service.UserService
	logger      *log.Logger
}

func NewHTTPHandler(userService *service.UserService, logger *log.Logger) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
		logger:      logger,
	}
}
