package routers

import (
	"github.com/gin-gonic/gin"
)

// Router const
const (
	VERSION = "/v1"
)

// RegisterRouters routes all api
func RegisterRouters(app *gin.Engine) {
	app.Group(VERSION)
	{

	}
}
