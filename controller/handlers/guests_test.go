package handlers

import (
	"github.com/benacook/getGround-technical-task/model/data"
	"github.com/benacook/getGround-technical-task/model/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

var (
	g = data.Guest{
		Id: 1,
		Name: "Ben",
		AdditionalGuests: 4,
		TableNumber: 1,
		Arrived: false,
	}
)
//======================================================================================
func TestGuestHandler_Put(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable()
	m.MockSprocGetVenue()
	m.MockSprocUpdateUsedCapacity(0)
	m.MockSprocGuestArrived()
	m.MockSprocGetGuestByName()

	req, err := http.NewRequest(http.MethodPut, "/guests/" + g.Name,
		strings.NewReader("{\"accompanying_guests\":" + strconv.Itoa(
			g.AdditionalGuests) + "}"))

	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	gh := NewGuestHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/guests").Subrouter()
	guestRouter.HandleFunc("/{name}", gh.Put).Methods(http.MethodPut)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusAccepted{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}
}

//======================================================================================
func TestGuestHandler_Put_Fail(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable()
	m.MockSprocGuestArrived_NoGuest()
	m.MockSprocGetGuestByName()

	req, err := http.NewRequest(http.MethodPut, "/guests/" + g.Name,
		strings.NewReader("{\"accompanying_guests\":" + strconv.Itoa(
			g.AdditionalGuests) + "}"))

	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	gh := NewGuestHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/guests").Subrouter()
	guestRouter.HandleFunc("/{name}", gh.Put).Methods(http.MethodPut)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

//======================================================================================
func TestGuestHandler_PutFailEmptyBody(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable()
	m.MockSprocGuestArrived()
	m.MockSprocGetGuestByName()

	req, err := http.NewRequest(http.MethodPut, "/guests/ben",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	gh := NewGuestHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/guests").Subrouter()
	guestRouter.HandleFunc("/{name}", gh.Put).Methods(http.MethodPut)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusInternalServerError{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

//======================================================================================
func TestGuestHandler_PutFailBadMethod(t *testing.T) {
	//no mocking needed as gorilla mux blocks this call

	req, err := http.NewRequest(http.MethodPut, "/guests",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	gh := NewGuestHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/guests").Subrouter()
	guestRouter.HandleFunc("/{name}", gh.Put).Methods(http.MethodPut)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusNotFound{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
//======================================================================================

func TestGuestHandler_Delete(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGuestLeft()

	req, err := http.NewRequest(http.MethodDelete, "/guests/Ben",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	gh := NewGuestHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/guests").Subrouter()
	guestRouter.HandleFunc("/{name}", gh.Delete).Methods(http.MethodDelete)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusAccepted{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}
}

//======================================================================================

func TestGuestHandler_Delete_Fail(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGuestLeft_NoGuest()

	req, err := http.NewRequest(http.MethodDelete, "/guests/Ben",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	gh := NewGuestHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/guests").Subrouter()
	guestRouter.HandleFunc("/{name}", gh.Delete).Methods(http.MethodDelete)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
//======================================================================================

func TestGuestHandler_Get(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetArrivedGuests()

	req, err := http.NewRequest(http.MethodGet, "/guests/ben",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	gh := NewGuestHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/guests").Subrouter()
	guestRouter.HandleFunc("/{name}", gh.Get).Methods(http.MethodGet)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//======================================================================================

func TestGuestHandler_Get_Fail(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetArrivedGuests_NoGuests()

	req, err := http.NewRequest(http.MethodGet, "/guests/ben",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	gh := NewGuestHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/guests").Subrouter()
	guestRouter.HandleFunc("/{name}", gh.Get).Methods(http.MethodGet)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusInternalServerError{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}



