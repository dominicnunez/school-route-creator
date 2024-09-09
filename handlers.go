package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addStudent(c *gin.Context) {
	var newStudent Student
	if err := c.BindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	students = append(students, newStudent)
	c.JSON(http.StatusCreated, newStudent)
}

func addSchool(c *gin.Context) {
	var newSchool School
	if err := c.BindJSON(&newSchool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	schools = append(schools, newSchool)
	c.JSON(http.StatusCreated, newSchool)
}

func generateRoutes(c *gin.Context) {
	routes := createRoutes(students, schools)
	c.JSON(http.StatusOK, routes)
}
