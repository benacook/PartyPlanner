package handlers

import (
	"github.com/benacook/GetGround-Assignment/model/data"
	"github.com/benacook/GetGround-Assignment/model/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

var (
	v = data.Venue{
		Id:             0,
		Name:           "London Hilton Bankside",
		Capacity:       200,
		NumberOfTables: 20,
		TableSize:      10,
		NextFreeTable:  1,
		UsedCapacity:   0,
	}
)

//======================================================================================
func TestGuestListHandler_Post(t *testing.T) {
	m:= database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable(g.TableNumber)
	m.MockSprocAddGuest()
	m.MockSprocGetGuestByName()
	m.MockSprocGetVenue()
	m.MockSprocUpdateUsedCapacity(5)

	req, err := http.NewRequest(http.MethodPost, "/guest_list/"+ g.Name,
	strings.NewReader("{\"accompanying_guests\":" + strconv.Itoa(
		g.AdditionalGuests) + ", \"tablenumber\":" + strconv.Itoa(g.TableNumber)+ "}"))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	glh := NewGuestListHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/guest_list").Subrouter()
	guestRouter.HandleFunc("/{name}", glh.Post).Methods(http.MethodPost)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusCreated{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

//======================================================================================
func TestGuestListHandler_Post_Fail(t *testing.T) {
	m:= database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable(g.TableNumber)
	m.MockSprocAddGuest_Error()
	m.MockSprocGetVenue()
	m.MockSprocUpdateUsedCapacity(1)

	req, err := http.NewRequest(http.MethodPost, "/guest_list/"+ g.Name,
		strings.NewReader("{\"accompanying_guests\":" + strconv.Itoa(
			g.AdditionalGuests) + ", \"tablenumber\":" + strconv.Itoa(g.TableNumber)+ "}"))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	glh := NewGuestListHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/guest_list").Subrouter()
	guestRouter.HandleFunc("/{name}", glh.Post).Methods(http.MethodPost)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

//======================================================================================
func TestGuestListHandler_Post_FailNoBody(t *testing.T) {

	req, err := http.NewRequest(http.MethodPost, "/guest_list/"+ g.Name,
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	glh := NewGuestListHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/guest_list").Subrouter()
	guestRouter.HandleFunc("/{name}", glh.Post).Methods(http.MethodPost)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusInternalServerError{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
//======================================================================================

func TestGuestListHandler_Delete(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocRemoveGuest()
	m.MockSprocGetVenue()
	m.MockSprocUpdateUsedCapacity( -5)

	req, err := http.NewRequest(http.MethodDelete, "/guest_list/Ben",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	glh := NewGuestListHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/guest_list").Subrouter()
	guestRouter.HandleFunc("/{name}", glh.Delete).Methods(http.MethodDelete)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusAccepted{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}
}

//======================================================================================

func TestGuestListHandler_Delete_Fail(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocRemoveGuest_Fail()

	req, err := http.NewRequest(http.MethodDelete, "/guest_list/Ben",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	glh := NewGuestListHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/guest_list").Subrouter()
	guestRouter.HandleFunc("/{name}", glh.Delete).Methods(http.MethodDelete)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
//======================================================================================

func TestGuestListHandler_Get(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetAllGuests()

	req, err := http.NewRequest(http.MethodGet, "/guest_list/Ben",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	glh := NewGuestListHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/guest_list").Subrouter()
	guestRouter.HandleFunc("/{name}", glh.Get).Methods(http.MethodGet)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//======================================================================================

func TestGuestListHandler_Get_Fail(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetAllGuests_NoGuests()

	req, err := http.NewRequest(http.MethodGet, "/guest_list/Ben",
		strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	rec := httptest.NewRecorder()

	glh := NewGuestListHandler()
	r := mux.NewRouter()
	guestRouter := r.PathPrefix("/guest_list").Subrouter()
	guestRouter.HandleFunc("/{name}", glh.Get).Methods(http.MethodGet)

	r.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusInternalServerError{
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}



