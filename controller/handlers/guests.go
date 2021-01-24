package handlers

import (
	"github.com/benacook/PartyPlanner-Assignment/model"
	"github.com/benacook/PartyPlanner-Assignment/model/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//======================================================================================

//GuestHandler type for class implements the interface httpHandler found in
//handler-interface.go.
type GuestHandler struct {

}

//======================================================================================

//NewGuestHandler returns a GuestHandler.
func NewGuestHandler() *GuestHandler {
	return &GuestHandler{}
}

//======================================================================================

//Get handles the get http requests to the /guests route.
//It returns a list of the arrived guests, if the list is empty,
//an empty list is returned, either way with status 200.
//Status 500 is returned if the call to get arrived guests fails.
func (h *GuestHandler) Get(w http.ResponseWriter, r *http.Request) {
	guests, err := model.GetArrivedGuests()
	if err != nil {
		log.Println(err)
		respondWithError(err, http.StatusInternalServerError, w)
		return
	}
	g := data.GuestList{Guests: guests}
	encodeResponseAsJSON(g, http.StatusOK, w)
}

//======================================================================================

//Put handles the put http requests to the /guests route.
//This route is used to mark a guest on the guest list as arrived and update their
//number of additional guests.
//It if successful returns the updated guest with status 200.
//Status 500 is returned if it fails to parse the user input.
//Status 400 is returned is marking the guest as arrived fails.
func (h *GuestHandler) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	g, err := data.ParseRequestGuest(r)

	if err != nil {
		log.Println(err)
		respondWithError(err, http.StatusInternalServerError, w)
		return
	}
	g.Name = name

	g, err = model.GuestArrival(g)
	if err != nil{
		log.Println(err)
		respondWithError(err, http.StatusBadRequest, w)
		return
	}
	encodeResponseAsJSON(g, http.StatusAccepted, w)
}

//======================================================================================

//Post is not used and is only here to implement the interface.
func (h *GuestHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

//======================================================================================

//Delete handles the delete http requests to the /guests route.
//It marks a user as no longer in attendance at the party.
//Status 202 is returned if the request was successful.
//Status 400 is returned if the call to mark the guest as leaving fails.
func (h *GuestHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	g := data.Guest{}
	g.Name = name

	if err := model.GuestLeaves(g); err != nil{
		log.Println(err)
		respondWithError(err, http.StatusBadRequest, w)
		return
	}
	respondWithMessage("Guest left the party", http.StatusAccepted, w)
}
