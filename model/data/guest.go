package data

import (
	"encoding/json"
	"net/http"
)

//======================================================================================

//Guest struct contains all of the data of a guest on the guest list,
//contains the same data as the guest_list table in the mysql database.
type Guest struct {
	Id int
	Name string
	AdditionalGuests int `json:"accompanying_guests"`
	TableNumber int
	Arrived bool
	ArrivalTime string
}

//======================================================================================

//GuestList contains a slice of Guest.
type GuestList struct {
	Guests []Guest
}

//======================================================================================

//ParseRequestGuest parses a http request to a guest struct.
//Returns a guest struct containing the parsed data and a nil error if successful.
//Returns a blank guest struct and an error if parsing fails.
func ParseRequestGuest(r *http.Request) (Guest, error) {
	dec := json.NewDecoder(r.Body)
	var g Guest
	err := dec.Decode(&g)
	if err != nil {
		return Guest{}, err
	}

	return g, nil
}