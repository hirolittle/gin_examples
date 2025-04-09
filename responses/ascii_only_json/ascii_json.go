package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON

func main() {
	router := gin.Default()

	router.GET("/someJson", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语言",
			"tag":  "<br>",
		}

		// Output: {"lang":"Go\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	router.Run(":8080")
}
