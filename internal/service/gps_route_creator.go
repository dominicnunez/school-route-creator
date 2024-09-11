// Package service contains the business logic for creating bus routes.
package service

import "route-creator/internal/model"

// createRoutes generates bus routes based on the provided students and schools
// It takes two slices as input: one for students and one for schools
// It returns a slice of Route structs
func CreateRoutes(students []model.Student, schools []model.School) []model.Route {
	// Initialize a slice of Route structs with the same length as the schools slice
	// This creates one route per school
	routes := make([]model.Route, len(schools))

	// Iterate over the schools slice
	// The range keyword is used for iteration in Go
	// 'i' is the index, and 'school' is the current School struct in the iteration
	for i, school := range schools {
		// Create a new Route for each school
		// The Students field is initialized as an empty slice
		routes[i] = model.Route{
			Students: []model.Student{},
			School:   school,
		}
	}

	// Iterate over the students slice
	for _, student := range students {
		// Check if there's at least one route
		if len(routes) > 0 {
			// Add the current student to the first route
			// This is a simplified approach and may not be ideal for real-world scenarios
			routes[0].Students = append(routes[0].Students, student)
		}
	}

	// Return the slice of routes
	return routes
}
