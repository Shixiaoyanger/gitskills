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

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
			"message": message,
			"nick":    nick,
		})
	})

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run(":8080")
}

func PrintHelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "hello,world")
}
