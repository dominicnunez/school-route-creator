// Package main contains the main application logic
package main

// createRoutes generates bus routes based on the provided students and schools
// It takes two slices as input: one for students and one for schools
// It returns a slice of Route structs
func createRoutes(students []Student, schools []School) []Route {
	// Initialize a slice of Route structs with the same length as the schools slice
	// This creates one route per school
	routes := make([]Route, len(schools))

	// Iterate over the schools slice
	// The range keyword is used for iteration in Go
	// 'i' is the index, and 'school' is the current School struct in the iteration
	for i, school := range schools {
		// Create a new Route for each school
		// The Students field is initialized as an empty slice
		routes[i] = Route{
			Students: []Student{},
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
