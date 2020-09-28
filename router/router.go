package router

import (
	"github.com/gin-gonic/gin"
	"go_apis/apis"
	"net/http"
)

func Routers() *gin.Engine{
	router := gin.Default()
	router.Use(cors())

	account := router.Group("/")
	{
		account.POST("/create", apis.CreateAccount)
		account.GET("/get", apis.GetAccount)
		account.POST("/login", apis.Login)
		account.GET("/getAll", apis.GetAccountAll)
	}

	return router
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-www-form-urlencoded")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
