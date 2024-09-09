// Package main contains the main application logic
package main

import (
	"net/http" // Provides HTTP client and server implementations

	"github.com/gin-gonic/gin" // Gin web framework for building web applications
)

// addStudent handles the addition of a new student
// It's called when a POST request is made to /students
func addStudent(c *gin.Context) {
	// Create a new Student struct to hold the incoming data
	var newStudent Student

	// Attempt to bind the JSON body of the request to our Student struct
	// If there's an error in binding (e.g., invalid JSON), return a 400 Bad Request error
	if err := c.BindJSON(&newStudent); err != nil {
		// Return a JSON response with the error message
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new student to our slice of students
	// Note: 'students' is likely defined in another file (e.g., models.go)
	students = append(students, newStudent)

	// Return a 201 Created status with the newly created student
	c.JSON(http.StatusCreated, newStudent)
}

// addSchool handles the addition of a new school
// It's called when a POST request is made to /schools
func addSchool(c *gin.Context) {
	// Create a new School struct to hold the incoming data
	var newSchool School

	// Attempt to bind the JSON body of the request to our School struct
	if err := c.BindJSON(&newSchool); err != nil {
		// Return a JSON response with the error message if binding fails
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new school to our slice of schools
	// Note: 'schools' is likely defined in another file (e.g., models.go)
	schools = append(schools, newSchool)

	// Return a 201 Created status with the newly created school
	c.JSON(http.StatusCreated, newSchool)
}

// generateRoutes creates bus routes based on the students and schools we have
// It's called when a GET request is made to /routes
func generateRoutes(c *gin.Context) {
	// Call the createRoutes function to generate routes
	// This function is likely defined in another file
	routes := createRoutes(students, schools)

	// Return the generated routes with a 200 OK status
	c.JSON(http.StatusOK, routes)
}
