package model

// Leg struct represents a single leg of a route
type Leg struct {
	ID       int    `json:"id"`
	SchoolID int    `json:"school_id"`
	Students []int  `json:"student_ids"`
	Stops    []Stop `json:"stops"`
	BusType  int    `json:"bus_type"`
}
