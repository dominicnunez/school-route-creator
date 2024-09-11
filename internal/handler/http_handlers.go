// Package handler contains HTTP request handlers for the route creation application.
package handler

import (
	"net/http"
	"route-creator/internal/data"    // Import your data package
	"route-creator/internal/model"   // Import your model package
	"route-creator/internal/service" // Import your service package

	// Add any other necessary imports

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DataStore data.DataStore
}

func NewHandler(ds data.DataStore) *Handler {
	return &Handler{DataStore: ds}
}

func (h *Handler) AddStudent(c *gin.Context) {
	var newStudent model.Student

	if err := c.BindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.DataStore.AddStudent(newStudent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newStudent)
}

func (h *Handler) AddSchool(c *gin.Context) {
	var newSchool model.School

	if err := c.BindJSON(&newSchool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.DataStore.AddSchool(newSchool)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newSchool)
}

func (h *Handler) GenerateRoutes(c *gin.Context) {
	students, err := h.DataStore.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve students"})
		return
	}

	schools, err := h.DataStore.GetSchools()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve schools"})
		return
	}

	routes := service.CreateRoutes(students, schools)
	c.JSON(http.StatusOK, routes)
}
