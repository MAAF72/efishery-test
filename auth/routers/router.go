package routers

import (
	"github.com/MAAF72/efishery-test/routers/auth"
	"github.com/gin-gonic/gin"
)

// Router const
const (
	VERSION = "/v1"
)

// RegisterRouters routes all api
func RegisterRouters(app *gin.Engine) {
	r := app.Group(VERSION)
	{
		auth.Routes(r)
	}
}
