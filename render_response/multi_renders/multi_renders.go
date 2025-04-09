package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

// json/xml/yaml/protobuf 渲染

func main() {
	router := gin.Default()

	// gin.H 是 map[string]interface{} 的一种快捷方式
	router.GET("/someJson", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello, json",
			"status":  http.StatusOK,
		})
	})

	router.GET("/moreJson", func(c *gin.Context) {
		// 你也可以使用一个结构体
		type Msg struct {
			Name    string `json:"user"`
			Message string `json:"message"`
			Number  int
		}
		var msg Msg
		msg.Name = "hiro"
		msg.Message = "hello, json"
		msg.Number = 123
		// 注意 msg.Name 在 JSON 中变成了 "user"
		// Message 字段，添加了 json tag，变成了 message
		// Number 字段，没有添加 json tag，Go 会自动用字段名（大小写不变）当作 JSON 的 key
		// Output: {"user":"hiro","message":"hello, json","Number":123}
		c.JSON(http.StatusOK, msg)
	})

	router.GET("/someXml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "hello, xml",
			"status":  http.StatusOK,
		})
	})

	router.GET("/someYaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"user": gin.H{
				"name":    "hiro",
				"age":     18,
				"hobbies": []string{"swimming", "running"},
			},
			"message": "hello, yaml",
			"status":  http.StatusOK,
		})
	})

	router.GET("/someProtobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})

	router.Run(":8080")
}
