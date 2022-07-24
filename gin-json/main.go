package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type msg struct {
	Name    string `json:"xingming"` // 结构体别名
	Message string
	Age     int
}

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		// data := map[string]interface{}{
		// 	"name":   "小明",
		// 	"age":    "8",
		// 	"gender": "男",
		// }

		data := gin.H{
			"name":   "小明",
			"age":    "8",
			"gender": "男",
		}

		c.JSON(http.StatusOK, data)
	})

	r.GET("/msg", func(c *gin.Context) {
		data := msg{
			Name:    "小红",
			Age:     17,
			Message: "JSON---ooo",
		}

		c.JSON(http.StatusOK, data)
	})

	r.Run(":9292")
}
