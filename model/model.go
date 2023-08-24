package model

import "github.com/go-playground/validator/v10"

type UserData struct {
	Username string `json:"username" validate:"required,min=2,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Age      string `json:"age" validate:"required,gte=13,lte=99"`
}

var validate = validator.New()
