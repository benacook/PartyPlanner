package data

import (
	"encoding/json"
	"net/http"
)

//======================================================================================

//Venue struct contains all of the data of the venue,
//contains the same data as the venue table in the mysql database.
type Venue struct {
	Id int
	Name string
	Capacity int
	NumberOfTables int
	TableSize int
	NextFreeTable int
	UsedCapacity int
}

//======================================================================================

//EmptySeats struct is used to return empty seats,
//either bookable or empty at the party.
type EmptySeats struct {
	SeatsEmpty int `json:"seats_empty"`
}

//======================================================================================

//ParseRequestVenue parses a http request to a venue struct.
//Returns a venue struct containing the parsed data and a nil error if successful.
//Returns a blank venue struct and an error if parsing fails.
func ParseRequestVenue(r *http.Request) (Venue, error) {
	dec := json.NewDecoder(r.Body)
	var g Venue
	err := dec.Decode(&g)
	if err != nil {
		return Venue{}, err
	}
	return g, nil
}
