// Package model defines the Address struct and related types.
package model

// Address struct represents a physical address and associated students by ID
// The `json:` tags specify how these fields should be named in JSON
type Address struct {
	StreetNumber int    `json:"street_number"`
	StreetPrefix string `json:"street_prefix"`
	StreetName   string `json:"street_name"`
	StreetSuffix string `json:"street_suffix"`
	StreetType   string `json:"street_type"`
	AddressUnit  string `json:"address_unit,omitempty"` // Optional field
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      int    `json:"zip_code"`
	Students     []int  `json:"students"`
}
