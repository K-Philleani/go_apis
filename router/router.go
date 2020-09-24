package router

import (
	"github.com/gin-gonic/gin"
	"go_apis/apis"
)

func Routers() *gin.Engine{
	router := gin.Default()

	account := router.Group("/")
	{
		account.POST("/create", apis.CreateAccount)
		account.GET("/get", apis.GetAccount)
		account.POST("/login", apis.Login)
	}

	return router
}
