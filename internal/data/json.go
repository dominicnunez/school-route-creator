// Package data provides a JSON file-based implementation of the DataStore interface for development.
package data

import (
	"encoding/json"
	"os"
	"route-creator/internal/model"
)

// JSONStore implements DataStore using JSON files
// This is just a skeleton, you'll need to implement the actual JSON reading/writing
type JSONStore struct {
	StudentsFile string // Path to the JSON file for students
	SchoolsFile  string // Path to the JSON file for schools
}

// NewJSONStore is a constructor function for JSONStore
func NewJSONStore(studentsFile, schoolsFile string) *JSONStore {
	return &JSONStore{
		StudentsFile: studentsFile,
		SchoolsFile:  schoolsFile,
	}
}

// AddStudent adds a new student to the JSONStore
func (js *JSONStore) AddStudent(student model.Student) error {
	// Read existing students from file
	students, err := js.GetStudents()
	if err != nil {
		return err
	}

	// Add new student
	students = append(students, student)

	// Write updated students back to file
	file, err := os.Create(js.StudentsFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(students)
}

// Implement other JSONStore methods here...
func (js *JSONStore) AddSchool(school model.School) error {
	schools, err := js.GetSchools()
	if err != nil {
		return err
	}
	schools = append(schools, school)
	file, err := os.Create(js.SchoolsFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(schools)
}

func (js *JSONStore) GetStudents() ([]model.Student, error) {
	file, err := os.Open(js.StudentsFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var students []model.Student
	err = json.NewDecoder(file).Decode(&students)
	return students, err
}

func (js *JSONStore) GetSchools() ([]model.School, error) {
	file, err := os.Open(js.SchoolsFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var schools []model.School
	err = json.NewDecoder(file).Decode(&schools)
	return schools, err
}
