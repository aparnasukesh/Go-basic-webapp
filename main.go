package main

import (
	"aparnasukesh/github.com/Go-basic-webapp/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.Routes(r)
	r.Run(":2000")
}
