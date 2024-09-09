package main

func createRoutes(students []Student, schools []School) []Route {
	routes := make([]Route, len(schools))
	for i, school := range schools {
		routes[i] = Route{
			Students: []Student{},
			School:   school,
		}
	}

	for _, student := range students {
		if len(routes) > 0 {
			routes[0].Students = append(routes[0].Students, student)
		}
	}

	return routes
}
