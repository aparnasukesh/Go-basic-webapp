package model

type UserData struct {
	Username string `json:"username" validate:"required,min=2,max=10"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Age      string `json:"age" validate:"required,gte=13,lte=99"`
}
