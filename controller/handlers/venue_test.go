package handlers

import (
	"github.com/benacook/PartyPlanner-Assignment/model/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)
//======================================================================================
func TestVenueHandler_Post(t *testing.T) {
	m := database.NewMock()
	m.MockSprocAddVenue()
	m.MockSprocGetAllGuests()
	m.MockSprocUpdateUsedCapacity(5)
	m.MockSprocGetVenue()

	req, err := http.NewRequest(http.MethodPost, "/venue",
		strings.NewReader("{\"name\": \"" + v.Name + "\", " +
			"\"capacity\": "+ strconv.Itoa(v.Capacity) + ", " +
			"\"numberoftables\": "+ strconv.Itoa(v.NumberOfTables) + "}"))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	vh := NewVenueHandler()
	r := mux.NewRouter()
	venueRouter := r.PathPrefix("/venue").Subrouter()
	venueRouter.HandleFunc("", vh.Post).Methods(http.MethodPost)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusCreated{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

//======================================================================================
func TestVenueHandler_Post_Fail(t *testing.T) {
	m := database.NewMock()
	m.MockSprocAddVenue_NoVenue()

	req, err := http.NewRequest(http.MethodPost, "/venue",
		strings.NewReader("{\"name\": \"" + v.Name + "\", " +
			"\"capacity\": "+ strconv.Itoa(v.Capacity) + ", " +
			"\"numberoftables\": "+ strconv.Itoa(v.NumberOfTables) + "}"))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	vh := NewVenueHandler()
	r := mux.NewRouter()
	venueRouter := r.PathPrefix("/venue").Subrouter()
	venueRouter.HandleFunc("", vh.Post).Methods(http.MethodPost)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

//======================================================================================
func TestVenueHandler_Post_FailBadRequest(t *testing.T) {
	m := database.NewMock()
	m.MockSprocAddVenue()
	m.MockSprocGetAllGuests()
	m.MockSprocUpdateUsedCapacity(5)
	m.MockSprocGetVenue()

	req, err := http.NewRequest(http.MethodPost, "/venue",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	vh := NewVenueHandler()
	r := mux.NewRouter()
	venueRouter := r.PathPrefix("/venue").Subrouter()
	venueRouter.HandleFunc("", vh.Post).Methods(http.MethodPost)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusInternalServerError{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

//======================================================================================

func TestGVenueHandler_Delete(t *testing.T) {
	m := database.NewMock()
	m.MockSprocDeleteVenue()

	req, err := http.NewRequest(http.MethodDelete, "/venue",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	vh := NewVenueHandler()
	r := mux.NewRouter()
	venueRouter := r.PathPrefix("/venue").Subrouter()
	venueRouter.HandleFunc("", vh.Delete).Methods(http.MethodDelete)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//======================================================================================

func TestGVenueHandler_Delete_Fail(t *testing.T) {
	m := database.NewMock()
	m.MockSprocDeleteVenue_NoVenue()

	req, err := http.NewRequest(http.MethodDelete, "/venue",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	vh := NewVenueHandler()
	r := mux.NewRouter()
	venueRouter := r.PathPrefix("/venue").Subrouter()
	venueRouter.HandleFunc("", vh.Delete).Methods(http.MethodDelete)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusInternalServerError{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
//======================================================================================

func TestVenueHandler_Get(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()

	req, err := http.NewRequest(http.MethodGet, "/venue",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	vh := NewVenueHandler()
	r := mux.NewRouter()
	venueRouter := r.PathPrefix("/venue").Subrouter()
	venueRouter.HandleFunc("", vh.Get).Methods(http.MethodGet)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//======================================================================================

func TestVenueHandler_Get_Fail(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue_NoVenue()

	req, err := http.NewRequest(http.MethodGet, "/venue",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	vh := NewVenueHandler()
	r := mux.NewRouter()
	venueRouter := r.PathPrefix("/venue").Subrouter()
	venueRouter.HandleFunc("", vh.Get).Methods(http.MethodGet)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusInternalServerError{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

//======================================================================================

func TestVenueHandler_GetSeats(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocGetArrivedGuests()

	req, err := http.NewRequest(http.MethodGet, "/seats_empty",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	vh := NewVenueHandler()
	r := mux.NewRouter()
	venueRouter := r.PathPrefix("/seats_empty").Subrouter()
	venueRouter.HandleFunc("", vh.GetSeats).Methods(http.MethodGet)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//======================================================================================

func TestVenueHandler_GetSeats_Fail(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue_NoVenue()

	req, err := http.NewRequest(http.MethodGet, "/seats_empty",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	vh := NewVenueHandler()
	r := mux.NewRouter()
	venueRouter := r.PathPrefix("/seats_empty").Subrouter()
	venueRouter.HandleFunc("", vh.GetSeats).Methods(http.MethodGet)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusInternalServerError{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

//======================================================================================

func TestVenueHandler_GetBookableSeats(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()

	req, err := http.NewRequest(http.MethodGet, "/seats_bookable",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	vh := NewVenueHandler()
	r := mux.NewRouter()
	venueRouter := r.PathPrefix("/seats_bookable").Subrouter()
	venueRouter.HandleFunc("", vh.GetBookableSeats).Methods(http.MethodGet)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//======================================================================================

func TestVenueHandler_GetBookableSeats_Fail(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue_NoVenue()

	req, err := http.NewRequest(http.MethodGet, "/seats_bookable",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	vh := NewVenueHandler()
	r := mux.NewRouter()
	venueRouter := r.PathPrefix("/seats_bookable").Subrouter()
	venueRouter.HandleFunc("", vh.GetBookableSeats).Methods(http.MethodGet)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusInternalServerError{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}



