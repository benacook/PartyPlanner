package handlers

import (
	"github.com/benacook/GetGround-Assignment/model"
	"github.com/benacook/GetGround-Assignment/model/data"
	"log"
	"net/http"
)

//======================================================================================

//GuestListHandler type for class implements the interface httpHandler found in
//handler-interface.go.
type VenueHandler struct {

}

//======================================================================================

//NewVenueHandler returns a VenueHandler.
func NewVenueHandler() *VenueHandler {
	return &VenueHandler{}
}

//======================================================================================

//Get handles the get http requests to the /venue route.
//It returns the venue data with status 200 if successful.
//Status 500 is returned if the call to get the venue fails.
func (h *VenueHandler) Get(w http.ResponseWriter, r *http.Request) {
	v, err := model.GetVenue()
	if err != nil {
		log.Println(err)
		respondWithError(err, http.StatusInternalServerError, w)
		return
	}
	encodeResponseAsJSON(v, http.StatusOK, w)
}

//======================================================================================

//GetSeats handles the get http requests to the /seats_empty route.
//It returns the number of empty seats during the party with status 200 if successful.
//Status 500 is returned if the call to get the remaining seats fails.
func (h *VenueHandler) GetSeats(w http.ResponseWriter, r *http.Request) {
	numSeats, err := model.GetRemainingSeats()
	if err != nil {
		log.Println(err)
		respondWithError(err, http.StatusInternalServerError, w)
		return
	}
	es := data.EmptySeats{SeatsEmpty: numSeats}
	encodeResponseAsJSON(es, http.StatusOK, w)
}

//======================================================================================

//GetBookableSeats handles the get http requests to the /seats_bookable route.
//It returns the number of un-allocated seats that can be used to invite guests with
//status 200 if successful.
//Status 500 is returned if the call to get the bookable seats fails.
func (h *VenueHandler) GetBookableSeats(w http.ResponseWriter, r *http.Request) {
	numSeats, err := model.GetRemainingBookableSeats()
	if err != nil {
		log.Println(err)
		respondWithError(err, http.StatusInternalServerError, w)
		return
	}
	es := data.EmptySeats{SeatsEmpty: numSeats}
	encodeResponseAsJSON(es, http.StatusOK, w)
}

//======================================================================================

//Put is not implemented and is defined to fulfill the interface.
func (h *VenueHandler) Put(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

//======================================================================================

//Post handles the post http requests to the /venue route.
//This route is used to add the venue to the list.
//If the call to add the venue is successful,
//the venue data is returned as a json object with status code 201.
//Status 500 is returned for errors parsing the user data.
//Status 400 is returned if adding the venue fails.
func (h *VenueHandler) Post(w http.ResponseWriter, r *http.Request) {
	v, err := data.ParseRequestVenue(r)
	if err != nil {
		log.Println(err)
		respondWithError(err, http.StatusInternalServerError, w)
		return
	}

	newVenue, err := model.AddVenue(v)
	if err != nil {
		log.Println(err)
		respondWithError(err, http.StatusBadRequest, w)
		return
	}
	encodeResponseAsJSON(newVenue, http.StatusCreated, w)
}

//======================================================================================

//Delete handles the delete http requests to the /venue route.
//This route is used to remove the venue.
//If the call to delete the venue is successful,
//Status 202 is returned if the call is successful.
//Status 500 is returned if the delete fails.
func (h *VenueHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := model.DeleteVenue(); err != nil {
		respondWithError(err, http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

