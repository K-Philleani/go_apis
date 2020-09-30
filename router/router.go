package router

import (
	"github.com/gin-gonic/gin"
	"go_apis/apis"
	"go_apis/middleware"
)

func Routers() *gin.Engine{
	router := gin.Default()
	router.Use(middleware.Cors())

	account := router.Group("/")
	{
		account.POST("/create", apis.CreateAccount)
		account.GET("/get", apis.GetAccount)
		account.POST("/login", apis.Login)
		account.GET("/getAll", apis.GetAccountAll)
		account.GET("/delete", apis.DeleteAccount)
	}

	return router
}




