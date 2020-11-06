package handlers

import (
	"encoding/json"
	"github.com/benacook/getGround-technical-task/model/data"
	"io"
	"net/http"
)

//======================================================================================

//encodeResponseAsJSON takes a blank interface,
//converts it to json and writes it to the io.writer w.
func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

//======================================================================================

//ParseRequestGuest parses a http request to a guest struct.
//Returns a guest struct containing the parsed data and a nil error if successful.
//Returns a blank guest struct and an error if parsing fails.
func ParseRequestGuest(r *http.Request) (data.Guest, error) {
	dec := json.NewDecoder(r.Body)
	var g data.Guest
	err := dec.Decode(&g)
	if err != nil {
		return data.Guest{}, err
	}
	return g, nil
}

//======================================================================================

//ParseRequestVenue parses a http request to a venue struct.
//Returns a venue struct containing the parsed data and a nil error if successful.
//Returns a blank venue struct and an error if parsing fails.
func ParseRequestVenue(r *http.Request) (data.Venue, error) {
	dec := json.NewDecoder(r.Body)
	var g data.Venue
	err := dec.Decode(&g)
	if err != nil {
		return data.Venue{}, err
	}
	return g, nil
}
