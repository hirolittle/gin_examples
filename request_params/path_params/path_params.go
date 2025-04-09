package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 此 handler 将匹配 /user/hiro，不会匹配 /user/ 和 /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 此 handler 将匹配 /user/hiro/send 和 /user/hiro/
	// 如果没有其他匹配 /user/hiro 将重定向到 /user/hiro/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, "Hello %s, %s", name, action)
	})

	router.Run(":8080")
}
