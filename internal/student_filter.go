package service

import (
	"route-creator/internal/model"
)

func getStudentsBySchool(students []model.Student, school model.School) []model.Student {
	var studentsBySchool []model.Student
	for _, student := range students {
		if student.SchoolID == school.ID {
			studentsBySchool = append(studentsBySchool, student)
		}
	}
	return studentsBySchool
}
