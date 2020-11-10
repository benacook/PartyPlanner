package handlers

import (
	"github.com/benacook/GetGround-Assignment/model"
	"github.com/benacook/GetGround-Assignment/model/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//======================================================================================

//GuestListHandler type for class implements the interface httpHandler found in
//handler-interface.go.
type GuestListHandler struct {

}

//======================================================================================

//NewGuestListHandler returns a GuestListHandler.
func NewGuestListHandler() *GuestListHandler {
	return &GuestListHandler{}
}

//======================================================================================

//Get handles the get http requests to the /guest_list route.
//It returns the current guest list, if the list is empty,
//an empty list is returned, either way with status 200.
//Status 500 is returned if the call to get all guests fails.
func (h *GuestListHandler) Get(w http.ResponseWriter, r *http.Request) {
	guests, err := model.GetAllGuests()
	if err != nil {
		log.Println(err)
		respondWithError(err, http.StatusInternalServerError, w)
		return
	}

	gl := data.GuestList{guests}
	encodeResponseAsJSON(gl, http.StatusOK, w)
}

//======================================================================================

//Put is not implemented and is defined to fulfill the interface.
func (h *GuestListHandler) Put(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

//======================================================================================

//Post handles the post http requests to the /guest_list route.
//This route is used to add new guests to the list.
//If the call to add the guest is successful,
//the guest is returned as a json object with status code 201.
//Status 500 is returned for errors parsing the user data.
//Status 400 is returned if adding the user fails.
func (h *GuestListHandler) Post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	g, err := data.ParseRequestGuest(r)
	if err != nil {
		log.Println(err)
		respondWithError(err, http.StatusInternalServerError, w)
		return
	}
	g.Name = name

	result, err := model.AddGuest(g)
	if err != nil{
		log.Println(err)
		respondWithError(err, http.StatusBadRequest, w)
		return
	}
	encodeResponseAsJSON(result, http.StatusCreated, w)
}

//======================================================================================

//Delete handles the delete http requests to the /guest_list route.
//This route is used to remove guests from the list.
//If the call to delete the guest is successful,
//Status 202 is returned if the cal is successful.
//Status 400 is returned if the delete fails.
func (h *GuestListHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	g := data.Guest{}
	g.Name = name

	if err := model.DeleteGuest(g); err != nil{
		log.Println(err)
		respondWithError(err, http.StatusBadRequest, w)
		return
	}
	respondWithMessage("Guest removed", http.StatusAccepted, w)
}
