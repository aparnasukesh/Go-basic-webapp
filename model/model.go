package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type UserData struct {
	Username string `json:"username" validate:"required,min=2,max=10"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,max=10"`
	Age      int    `json:"age" validate:"required,gte=13,lte=99"`
}

func Validate(userData UserData) error {
	validate := validator.New()
	err := validate.Struct(userData)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, e := range validationErrors {

			if e.Field() == "Email" {
				return fmt.Errorf("The Email Address You Entered is Invalid")
			} else if e.Field() == "Age" {
				return fmt.Errorf("Invalid Age, Age should be between 13 to 99")
			} else if e.Field() == "Password" {
				return fmt.Errorf("Invalid Password. Password should be at least 4 characters long")
			} else if e.Field() == "Username" {
				return fmt.Errorf("Invalid Username, user name should be minimum two charactore")
			}
		}

	}
	return nil
}
