package main

import "time"

type Student struct {
	Name         string `json:"name"`
	Address      string `json:"address"`
	SpecialNeeds string `json:"special_needs"`
}

type School struct {
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type Route struct {
	Students []Student `json:"students"`
	School   School    `json:"school"`
}

var students []Student
var schools []School
