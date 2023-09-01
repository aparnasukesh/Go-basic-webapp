package middlware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if session.Get("username") == nil {
			ctx.Redirect(http.StatusSeeOther, "/login")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func Middleware(ctx *gin.Context) {
	session := sessions.Default(ctx)
	if session.Get("username") != nil {
		ctx.Redirect(http.StatusSeeOther, "/home")
		ctx.Abort()
		return
	}

	ctx.Next()
}
