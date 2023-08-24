package controller

import (
	"html/template"
	"net/http"

	"aparnasukesh/github.com/Go-basic-webapp/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var signupData = make(map[string]model.UserData)

func Routes(r *gin.Engine) {

	r.GET("/login", loginPage)
	r.GET("/signup", signUpPage)
	r.GET("/home", homePage)

	r.POST("/login", Login)
	r.POST("/signup", signUp)
	r.POST("/logout", Logout)
}

func loginPage(ctx *gin.Context) {
	temp, err := template.ParseFiles("view/login.html")
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": err,
		})
	}
	err = temp.Execute(ctx.Writer, nil)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": err,
		})
	}
}

func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if user, found := signupData[username]; found && user.Password == password {
		session := sessions.Default(ctx)
		session.Set("username", username)
		session.Save()
		ctx.Redirect(http.StatusMovedPermanently, "/home")
	} else {
		ctx.JSON(400, gin.H{
			"Message": "invalide username and password",
		})
	}
}

func signUpPage(ctx *gin.Context) {
	temp, err := template.ParseFiles("view/signup.html")
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": err,
		})
	}
	err = temp.Execute(ctx.Writer, nil)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": err,
		})
	}
}

func signUp(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	age := ctx.PostForm("age")

	if username == "" || password == "" || email == "" || age == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	userdata := model.UserData{
		Username: username,
		Password: password,
		Email:    email,
		Age:      age,
	}
	signupData[username] = userdata

	ctx.Redirect(http.StatusSeeOther, "/login")
}

func homePage(ctx *gin.Context) {
	temp, err := template.ParseFiles("view/home.html")
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": err,
		})
	}
	err = temp.Execute(ctx.Writer, nil)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": err,
		})
	}

}

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()

	ctx.Redirect(http.StatusSeeOther, "/login")
}
