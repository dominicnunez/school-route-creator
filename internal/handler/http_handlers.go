package handler

import (
	"fmt"
	"net/http"
	"route-creator/internal/data"
	"route-creator/internal/model"
	"route-creator/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DataStore data.DataStore
}

func NewHandler(ds data.DataStore) *Handler {
	return &Handler{
		DataStore: ds,
	}
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

	var allLegs []model.Leg

	for _, school := range schools {

		amLegs, err := service.CreateLegs(school, students, model.AM)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create AM legs for school %d: %v", school.ID, err)})
			return
		}
		pmLegs, err := service.CreateLegs(school, students, model.PM)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create PM legs for school %d: %v", school.ID, err)})
			return
		}

		allLegs = append(allLegs, amLegs...)
		allLegs = append(allLegs, pmLegs...)
	}

	// TODO: Implement route creation logic
	// For now, just return the legs
	c.JSON(http.StatusOK, gin.H{"legs": allLegs})
}
