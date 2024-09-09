// Package main is the entry point of our application
package main

// Import the Gin web framework
import (
	"github.com/gin-gonic/gin"
)

// main function is the entry point of our Go program
func main() {
	// Create a default Gin router
	// This sets up a new Gin engine with the default middleware
	r := gin.Default()

	// Serve static files from the "static" directory
	// This allows us to serve files like CSS, JavaScript, and images
	// These files will be accessible under the "/static" URL path
	r.Static("/static", "./static")

	// Define a route for the root URL "/"
	// This will serve our index.html file when someone visits the homepage
	r.GET("/", func(c *gin.Context) {
		// Serve the index.html file from the static directory
		c.File("./static/index.html")
	})

	// Define API routes for our application
	// Each route specifies an HTTP method, a path, and a handler function

	// POST /students: Add a new student
	// The addStudent function (defined elsewhere) will handle this request
	r.POST("/students", addStudent)

	// POST /schools: Add a new school
	// The addSchool function (defined elsewhere) will handle this request
	r.POST("/schools", addSchool)

	// GET /routes: Generate and return routes
	// The generateRoutes function (defined elsewhere) will handle this request
	r.GET("/routes", generateRoutes)

	// Start the server and listen on port 8080
	// If there's an error starting the server, it will be logged
	r.Run(":8080")
}
