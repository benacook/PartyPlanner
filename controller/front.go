package controller

import (
	"github.com/benacook/PartyPlanner-Assignment/controller/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

//======================================================================================

//RegisterHandlers resisters the routes and handlers setup in GetRouter.
func RegisterHandlers(){
	r := GetRouter()
	http.Handle("/", r)
}

//======================================================================================

//GetRouter sets up the API routing and returns a router mux.
func GetRouter() *mux.Router {
	gh := handlers.NewGuestHandler()
	glh := handlers.NewGuestListHandler()
	ih := handlers.NewInviteHandler()
	vh := handlers.NewVenueHandler()
	r := mux.NewRouter()

	guestRouter := r.PathPrefix("/guests").Subrouter()
	guestRouter.HandleFunc("/{name:[A-Za-z]+}", gh.Delete).Methods(http.MethodDelete)
	guestRouter.HandleFunc("/{name:[A-Za-z]+}", gh.Put).Methods(http.MethodPut)
	guestRouter.HandleFunc("", gh.Get).Methods(http.MethodGet)

	guestListRouter := r.PathPrefix("/guest_list").Subrouter()
	guestListRouter.HandleFunc("/{name:[A-Za-z]+}", glh.Delete).Methods(http.MethodDelete)
	guestListRouter.HandleFunc("/{name:[A-Za-z]+}", glh.Post).Methods(http.MethodPost)
	guestListRouter.HandleFunc("", glh.Get).Methods(http.MethodGet)

	inviteRouter := r.PathPrefix("/invitation").Subrouter()
	inviteRouter.HandleFunc("/{name:[A-Za-z]+}", ih.Get).Methods(http.MethodGet)

	venueRouter := r.PathPrefix("/venue").Subrouter()
	venueRouter.HandleFunc("", vh.Delete).Methods(http.MethodDelete)
	venueRouter.HandleFunc("", vh.Post).Methods(http.MethodPost)
	venueRouter.HandleFunc("", vh.Get).Methods(http.MethodGet)

	partyRouter := r.PathPrefix("/seats_empty").Subrouter()
	partyRouter.HandleFunc("", vh.GetSeats).Methods(http.MethodGet)

	bookableSeatsRouter := r.PathPrefix("/seats_bookable").Subrouter()
	bookableSeatsRouter.HandleFunc("", vh.GetBookableSeats).Methods(http.MethodGet)

	return r
}

