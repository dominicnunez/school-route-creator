// Package main is the entry point of our route creation application.
// It sets up the web server and initializes the necessary components.
package main

import (
	"route-creator/internal/config"
	"route-creator/internal/data"
	"route-creator/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		// Handle the error
		panic(err)
	}

	// Choose data store based on configuration
	var dataStore data.DataStore
	switch cfg.DataStore {
	case "memory":
		dataStore = data.NewMemoryStore()
	case "json":
		dataStore = data.NewJSONStore("students.json", "schools.json")
	default:
		// Handle invalid data store type
		panic("Invalid data store type")
	}

	// Create a new Handler instance
	h := handler.NewHandler(dataStore)

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
	r.POST("/students", h.AddStudent)

	// POST /schools: Add a new school
	// The addSchool function (defined elsewhere) will handle this request
	r.POST("/schools", h.AddSchool)

	// GET /routes: Generate and return routes
	// The generateRoutes function (defined elsewhere) will handle this request
	r.GET("/routes", h.GenerateRoutes)

	// Start the server and listen on port 8080
	// If there's an error starting the server, it will be logged
	r.Run(":8080")
}
