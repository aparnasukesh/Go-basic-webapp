package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"aparnasukesh/github.com/Go-basic-webapp/middlware"
	"aparnasukesh/github.com/Go-basic-webapp/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var signupData = make(map[string]model.UserData)

type validate struct {
	message string
}

func Routes(r *gin.Engine) {

	store := cookie.NewStore([]byte("xyz"))
	r.Use(sessions.Sessions("session", store))

	r.GET("/login", middlware.Middleware, loginPage)
	r.GET("/signup", middlware.Middleware, signUpPage)
	r.GET("/home", middlware.AuthRequired(), homePage)

	r.POST("/login", Login)
	r.POST("/signup", signUp)
	r.POST("/logout", Logout)
}

func loginPage(ctx *gin.Context) {

	ctx.Header("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")

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
	ctx.Done()
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if username == "" || password == "" {
		ctx.JSON(200, gin.H{"error": "All fields are required"})
		return
	}

	if user, found := signupData[username]; found && user.Password == password {
		session := sessions.Default(ctx)
		session.Set("username", username)
		session.Save()
		ctx.Redirect(http.StatusMovedPermanently, "/home")
	} else {
		ctx.JSON(200, gin.H{
			"error": "Invalide username and password",
		})
	}
}

func signUpPage(ctx *gin.Context) {

	ctx.Header("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")

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
	ageStr := ctx.PostForm("age")

	if username == "" || password == "" || email == "" || ageStr == "" {
		ctx.JSON(200, gin.H{"error": "All fields are required"})
		return
	}

	age, err := strconv.Atoi(ageStr)

	userdata := model.UserData{
		Username: username,
		Password: password,
		Email:    email,
		Age:      age,
	}

	err = model.Validate(userdata)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
		return
	}

	signupData[username] = userdata

	ctx.Redirect(http.StatusSeeOther, "/login")
}

func homePage(ctx *gin.Context) {
	session := sessions.Default(ctx)
	username := session.Get("username")

	ctx.Header("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")

	temp, err := template.ParseFiles("view/home.html")
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": err,
		})
		return
	}

	data := map[string]interface{}{
		"Username": username,
	}

	err = temp.Execute(ctx.Writer, data)
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
