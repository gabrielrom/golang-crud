package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example/crud/dtos"
	"example/crud/models"

	"github.com/google/uuid"
)

type UserController struct {
	users []models.User
}

// Contructor
func (user *UserController) Init() {
	user.users = []models.User{}
}

func (u *UserController) CreateAccount(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")

	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusForbidden)
		fmt.Fprint(res, "Bad Request")
		return
	}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	var userRequestDto dtos.UserRequestDto

	if err := decoder.Decode(&userRequestDto); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, "Failed to parse request")
		return
	}

	if userRequestDto.Name == "" ||
		userRequestDto.Email == "" ||
		userRequestDto.Password == "" {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, "Missing fields")
		return
	}

	user := models.User{
		Id:       uuid.New().String(),
		Name:     userRequestDto.Name,
		Email:    userRequestDto.Email,
		Password: userRequestDto.Password,
	}

	u.users = append(u.users, user)

	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, dtos.NewUserResponse(user))
}

func (u *UserController) ListAccounts(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusForbidden)
		fmt.Fprint(res, "Error")
		return
	}

	str, err := json.Marshal(u.users)

	if err != nil {
		res.WriteHeader(http.StatusForbidden)
		fmt.Fprint(res, "Error json")
		return
	}

	fmt.Fprint(res, string(str))
}
