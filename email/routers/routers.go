package routers

import (
	"email/controllers"
	"email/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() {

	router := gin.Default()

	router.POST("/regist", controllers.Regist)
	router.POST("/login", controllers.Login)

	tar := router.Group("/v1")
	tar.Use(jwt.JWT())
	{
		tar.GET("/test", PrintHelloWorld)
		tar.POST("/send", controllers.Send)
		tar.GET("/receive", controllers.Receive)
	}


	router.Run(":8080")
}

func PrintHelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "hello,world")
}
