package model

import "time"

type Stop struct {
	ID           int       `json:"id"`
	Intersection string    `json:"intersection"`
	Corner       string    `json:"corner"`
	Longitude    float64   `json:"longitude"`
	Latitude     float64   `json:"latitude"`
	AMPickup     time.Time `json:"am_pickup"`
	PMDropoff    time.Time `json:"pm_dropoff"`
}
