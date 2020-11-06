package model

import (
	"github.com/benacook/getGround-technical-task/model/data"
	"github.com/benacook/getGround-technical-task/model/database"
	"testing"
)

var (
	guest = data.Guest{
		Id:               0,
		Name:             "Ben",
		AdditionalGuests: 4,
		TableNumber:      1,
		Arrived:          false,
		ArrivalTime:      "",
	}
)

func TestAddGuest_FailNoVenue(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue_NoVenue()
	_, err := AddGuest(guest)
	if err == nil {
		t.Fatal(err)
	}
}

func TestAddGuest_FailCantGetGuests(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable(guest.TableNumber)
	for i := 1; i <= v.NumberOfTables; i++{
		m.MockSprocGetGuestsAtTable(i)
	}

	g := guest
	g.AdditionalGuests = 100
	_, err := AddGuest(g)
	if err == nil {
		t.Fatal(err)
	}
}

func TestAddGuest_FailNoSpace(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable_NoGuests(guest.TableNumber)
	for i := 1; i <= v.NumberOfTables; i++{
		m.MockSprocGetGuestsAtTable_NoGuests(guest.TableNumber)
	}
	_, err := AddGuest(guest)
	if err == nil {
		t.Fatal(err)
	}
}

func TestGetArrivedGuests_Fail(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetArrivedGuests_NoGuests()
	_, err := GetArrivedGuests()
	if err == nil {
		t.Fatal(err)
	}
}

func TestDeleteGuest_Fail(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName_NoGuest()
	err := DeleteGuest(guest)
	if err == nil {
		t.Fatal(err)
	}
}

func TestGuestArrival_FailNoGuest(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName_NoGuest()
	_, err := GuestArrival(guest)
	if err == nil {
		t.Fatal(err)
	}
}

func TestGuestArrival_FailNoVenue(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGetVenue_NoVenue()
	_, err := GuestArrival(guest)
	if err == nil {
		t.Fatal(err)
	}
}

func TestGuestArrival_FailNoSpace(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable(guest.TableNumber)
	g := guest
	g.AdditionalGuests = 100
	_, err := GuestArrival(g)
	if err == nil {
		t.Fatal(err)
	}
}

func TestGuestArrival_FailArrival(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable(guest.TableNumber)
	m.MockSprocGetVenue()
	m.MockSprocUpdateUsedCapacity(0)
	m.MockSprocGuestArrived_NoGuest()
	_, err := GuestArrival(guest)
	if err == nil {
		t.Fatal(err)
	}
}

func TestGuestArrival_FailGetUpdatedGuest(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable(guest.TableNumber)
	m.MockSprocGetVenue()
	m.MockSprocUpdateUsedCapacity(0)
	m.MockSprocGuestArrived()
	m.MockSprocGetGuestByName_NoGuest()
	_, err := GuestArrival(guest)
	if err == nil {
		t.Fatal(err)
	}
}

func TestGuestLeaves_FailNotOnList(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName_NoGuest()
	err := GuestLeaves(guest)
	if err == nil {
		t.Fatal(err)
	}
}

func TestGuestLeaves_FailToLeave(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGuestLeft_NoGuest()
	err := GuestLeaves(guest)
	if err == nil {
		t.Fatal(err)
	}
}