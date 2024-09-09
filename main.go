package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static files from the "static" directory
	r.Static("/static", "./static")

	// Serve index.html
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	r.POST("/students", addStudent)
	r.POST("/schools", addSchool)
	r.GET("/routes", generateRoutes)

	r.Run(":8080")
}
