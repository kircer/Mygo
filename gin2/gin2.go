package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 路由分组：/api/v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{"data": "users list v1"})
		})
		v1.GET("/products", func(c *gin.Context) {
			c.JSON(200, gin.H{"data": "products list v1"})
		})
	}

	// 路由分组：/api/v2，并应用专属中间件
	v2 := r.Group("/api/v2")
	v2.Use(func(c *gin.Context) {
		println("这是 v2 专属的中间件")
		c.Next()
	})
	{
		v2.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{"data": "users list v2"})
		})
	}

	r.Run(":8080")
}
