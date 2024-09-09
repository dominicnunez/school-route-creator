// Package main is used for all files in this project since it's a single executable application
package main

import (
	"encoding/json"
	"os"
)

// DataStore interface defines methods for data operations
// Interfaces in Go provide a way to specify the behavior of an object
type DataStore interface {
	AddStudent(student Student) error
	AddSchool(school School) error
	GetStudents() ([]Student, error)
	GetSchools() ([]School, error)
}

// MemoryStore implements DataStore using in-memory slices
// This is a struct that will hold our data in memory
type MemoryStore struct {
	students []Student // Slice of Student structs
	schools  []School  // Slice of School structs
}

// NewMemoryStore is a constructor function for MemoryStore
// It initializes and returns a pointer to a new MemoryStore
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		students: make([]Student, 0), // Initialize an empty slice for students
		schools:  make([]School, 0),  // Initialize an empty slice for schools
	}
}

// AddStudent adds a new student to the MemoryStore
// The (ms *MemoryStore) syntax defines this as a method on MemoryStore
func (ms *MemoryStore) AddStudent(student Student) error {
	ms.students = append(ms.students, student) // Append the new student to the slice
	return nil                                 // Return nil as there's no error
}

// AddSchool adds a new school to the MemoryStore
func (ms *MemoryStore) AddSchool(school School) error {
	ms.schools = append(ms.schools, school) // Append the new school to the slice
	return nil
}

// GetStudents returns all students from the MemoryStore
func (ms *MemoryStore) GetStudents() ([]Student, error) {
	return ms.students, nil // Return the slice of students
}

// GetSchools returns all schools from the MemoryStore
func (ms *MemoryStore) GetSchools() ([]School, error) {
	return ms.schools, nil // Return the slice of schools
}

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
func (js *JSONStore) AddStudent(student Student) error {
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
func (js *JSONStore) AddSchool(school School) error {
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

func (js *JSONStore) GetStudents() ([]Student, error) {
	file, err := os.Open(js.StudentsFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var students []Student
	err = json.NewDecoder(file).Decode(&students)
	return students, err
}

func (js *JSONStore) GetSchools() ([]School, error) {
	file, err := os.Open(js.SchoolsFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var schools []School
	err = json.NewDecoder(file).Decode(&schools)
	return schools, err
}

// Add a placeholder for DBStore (for future production use)
type DBStore struct {
	// Add database connection details here
}

func NewDBStore() *DBStore {
	// Initialize database connection
	return &DBStore{}
}

// Implement DBStore methods (placeholders)
func (db *DBStore) AddStudent(student Student) error {
	// TODO: Implement database insertion
	return nil
}

func (db *DBStore) AddSchool(school School) error {
	// TODO: Implement database insertion
	return nil
}

func (db *DBStore) GetStudents() ([]Student, error) {
	// TODO: Implement database query
	return nil, nil
}

func (db *DBStore) GetSchools() ([]School, error) {
	// TODO: Implement database query
	return nil, nil
}
