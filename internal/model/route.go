package model

// School struct represents a single school
// BusRoute struct represents a single bus route
type BusRoute struct {
	RouteID  int    `json:"route_id"`
	SchoolID int    `json:"school_id"`
	Students []int  `json:"student_ids"`
	Stops    []Stop `json:"stops"`
	BusType  int    `json:"bus_type"`
}
