// Package handler contains HTTP request handlers for the route creation application.
package handler

import (
    "net/http"
    "route-creator/internal/data"    // Import your data package
    "route-creator/internal/model"   // Import your model package
    "route-creator/internal/service" // Import your service package

    "github.com/gin-gonic/gin"
)

type Handler struct {
    DataStore data.DataStore
    LegCreator service.LegCreator
    RouteCreator service.RouteCreator
}

func NewHandler(ds data.DataStore, lc service.LegCreator, rc service.RouteCreator) *Handler {
    return &Handler{
        DataStore: ds,
        LegCreator: lc,
        RouteCreator: rc,
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
        // Get the students attending this school
        schoolStudents := getStudentsForSchool(students, school)

        // Create the legs for this school (both AM and PM)
        amLegs := h.LegCreator.CreateLegs(school, schoolStudents, service.AM)
        pmLegs := h.LegCreator.CreateLegs(school, schoolStudents, service.PM)

        // Append the legs to allLegs
        allLegs = append(allLegs, amLegs...)
        allLegs = append(allLegs, pmLegs...)
    }

    // Create routes from all legs
    routes := h.RouteCreator.CreateRoutes(allLegs)

    c.JSON(http.StatusOK, routes)
}

// getStudentsForSchool is a helper function to get the students attending a specific school
 {
    var schoolStudents []model.Student
    for _, student := range students {
        if student.SchoolID == school.ID {
            schoolStudents = append(schoolStudents, student)
        }
    }
    return schoolStudents
}
