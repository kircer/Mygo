package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义接收 JSON 的结构体
type Login struct {
	User     string `form:"user" json:"user" binding:"required"` // binding 标签用于校验
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	// GET 请求：获取 URL 路径参数和 Query 参数
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")               // 获取路径参数 /user/:name
		query := c.DefaultQuery("foo", "bar") // 获取 query 参数 /user/:name?foo=xxx
		c.String(http.StatusOK, "User %s, foo=%s", name, query)
	})

	// POST 请求：解析 JSON 请求体
	r.POST("/login", func(c *gin.Context) {
		var json Login
		// ShouldBindJSON 会根据 Content-Type 自动解析，并进行 binding 校验
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User == "admin" && json.Password == "123456" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	})

	r.Run(":8080")
}
