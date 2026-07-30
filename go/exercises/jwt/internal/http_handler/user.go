package http_handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gostudy/internal/model"
	"gostudy/pkg/utils"
	"io"
	"log"
	"net/http"
	"os"
)

func (h *UserHTTPHandler) GetProfile(r http.ResponseWriter, req *http.Request) {
	contex := req.Context()
	select {
	case <-contex.Done():
		return
	default:
		switch req.Method {
		case http.MethodGet:
			tokenString := req.Header.Get("Authorization")
			if tokenString == "" {
				r.WriteHeader(http.StatusUnauthorized)
				r.Header().Set("Allow", "POST")
				r.Header().Set("Content-Type", "application/json")
				_, err := r.Write([]byte(`{"message": "Authorization required"}`))
				if err != nil {
					log.Println(err.Error())
					return
				}
			}
			userProfile, err := h.userService.GetProfile(tokenString)
			if err != nil {
				r.WriteHeader(http.StatusUnauthorized)
				r.Header().Set("Allow", "POST")
				r.Header().Set("Content-Type", "application/json")
				_, err := r.Write([]byte(`{"message": "Invalid token"}`))
				log.Println(err.Error())
				return
			}
			err = json.NewEncoder(r).Encode(userProfile)
			if err != nil {
				r.WriteHeader(http.StatusInternalServerError)
				r.Header().Set("Allow", "POST")
				r.Header().Set("Content-Type", "application/json")
				h.logger.Println(err.Error())
				return
			}

		default:
			r.WriteHeader(http.StatusMethodNotAllowed)
			r.Header().Set("Allow", "POST")
			r.Header().Set("Content-Type", "application/json")
			_, err := r.Write([]byte(`{"error":"method not allowed"}`))
			if err != nil {
				log.Println(err.Error())
				return
			}
			return
		}
	}
}

func (h *UserHTTPHandler) RegisterHandler(r http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	select {
	case <-ctx.Done():
		return
	default:
		switch req.Method {
		case http.MethodPost:
			body := make([]byte, req.ContentLength)
			_, err := req.Body.Read(body)
			if err != nil && err != io.EOF {
				h.logger.Println(err.Error())
				return
			}
			user := model.User{}
			err = json.NewDecoder(bytes.NewReader(body)).Decode(&user)
			if err != nil {
				h.logger.Println(err.Error())
				return
			}
			err = h.userService.Register(user.Name, user.Password)
			if err != nil {
				h.logger.Println(err.Error())
				return
			}
			r.WriteHeader(http.StatusCreated)
			r.Header().Set("Content-Type", "application/json")
			_, err = r.Write([]byte(`{"message":"user registered"}`))
			if err != nil {
				h.logger.Println(err.Error())
				return
			}
			return

		default:
			r.WriteHeader(http.StatusMethodNotAllowed)
			r.Header().Set("Allow", "POST")
			r.Header().Set("Content-Type", "application/json")
			_, err := r.Write([]byte(`{"error":"method not allowed"}`))
			if err != nil {
				log.Println(err.Error())
				return
			}
		}

	}
}

func (h *UserHTTPHandler) LoginHandler(r http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	select {
	case <-ctx.Done():
		return
	default:
		switch req.Method {
		case http.MethodPost:
			body := make([]byte, req.ContentLength)
			_, err := req.Body.Read(body)
			if err != nil && err != io.EOF {
				r.WriteHeader(http.StatusInternalServerError)
				r.Header().Set("Allow", "POST")
				r.Header().Set("Content-Type", "application/json")
				_, err := r.Write([]byte(`{"message": "Internal server error"}`))
				if err != nil {
					return
				}
				return
			}
			user := model.User{}
			err = json.NewDecoder(bytes.NewReader(body)).Decode(&user)
			if err != nil {
				r.WriteHeader(http.StatusInternalServerError)
				r.Header().Set("Allow", "POST")
				r.Header().Set("Content-Type", "application/json")
				_, err := r.Write([]byte(`{"message": "Internal server error"}`))
				if err != nil {
					return
				}
				return
			}
			findUser, err := h.userService.Login(user.Name, user.Password)
			if err != nil {
				r.WriteHeader(http.StatusForbidden)
				r.Header().Set("Allow", "POST")
				r.Header().Set("Content-Type", "application/json")
				_, err := r.Write([]byte(`{"message": "Invalid username or password"}`))
				h.logger.Println(err.Error())
				if err != nil {
					h.logger.Fatalln(err.Error())
				}
				return
			}
			token, err := utils.GenerateToken(findUser.ID, os.Getenv("SECRET_KEY"))
			if err != nil {
				r.WriteHeader(http.StatusInternalServerError)
				r.Header().Set("Allow", "POST")
				r.Header().Set("Content-Type", "application/json")
				_, err := r.Write([]byte(`{"message": "Internal server error"}`))
				h.logger.Fatalln(err.Error())
			}
			r.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
			r.WriteHeader(http.StatusAccepted)
			r.Header().Set("Content-Type", "application/json")
			r.Write([]byte(fmt.Sprintf("{ \"token\": \"%s\", \"user\": \"%s\"}", token, findUser.Name)))
			return
		}
	}
}
