package dtos

import "example/crud/models"

type UserResponseDto struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserResponse(req models.User) UserResponseDto {
	return UserResponseDto{
		Id:    req.Id,
		Name:  req.Name,
		Email: req.Email,
	}
}
