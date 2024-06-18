package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thefran23/hayvnAPI/controllers"
)

func MessagesRoute(router *gin.Engine) {
	router.POST("/vineyard", func(c *gin.Context) {}, controllers.Message())

}
