package handlers

import (
	"github.com/benacook/getGround-technical-task/model/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInviteHandler_Get(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestByName()

	req, err := http.NewRequest(http.MethodGet, "/invitation/" + g.Name,
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	ih := NewInviteHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/invitation").Subrouter()
	guestRouter.HandleFunc("/{name}", ih.Get).Methods(http.MethodGet)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}


func TestInviteHandler_Get_FailVenue(t *testing.T) {
	m:= database.NewMock()
	m.MockSprocGetVenue_NoVenue()
	m.MockSprocGetGuestByName()

	req, err := http.NewRequest(http.MethodGet, "/invitation/" + g.Name,
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	ih := NewInviteHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/invitation").Subrouter()
	guestRouter.HandleFunc("/{name}", ih.Get).Methods(http.MethodGet)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusInternalServerError{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}

func TestInviteHandler_Get_FailGuest(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestByName_NoGuest()

	req, err := http.NewRequest(http.MethodGet, "/invitation/" + g.Name,
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	ih := NewInviteHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/invitation").Subrouter()
	guestRouter.HandleFunc("/{name}", ih.Get).Methods(http.MethodGet)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusInternalServerError{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}