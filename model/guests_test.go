package model

import (
	"github.com/benacook/GetGround-Assignment/model/data"
	"github.com/benacook/GetGround-Assignment/model/database"
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

func TestAddGuest(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable(guest.TableNumber)
	m.MockSprocAddGuest()
	m.MockSprocGetGuestByName()
	m.MockSprocGetVenue()
	m.MockSprocUpdateUsedCapacity(guest.AdditionalGuests + 1)
	_, err := AddGuest(guest)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddGuest_FailAdd(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable(guest.TableNumber)
	m.MockSprocAddGuest_Error()
	_, err := AddGuest(guest)
	if err == nil {
		t.Fatal(err)
	}
}

func TestAddGuest_FailNoVenue(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue_NoVenue()
	_, err := AddGuest(guest)
	if err == nil {
		t.Fatal(err)
	}
}

func TestAddGuest_FailFullVenue(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue_FullVenue()
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

func TestGetArrivedGuests(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetArrivedGuests()
	_, err := GetArrivedGuests()
	if err != nil {
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

func TestDeleteGuest(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocRemoveGuest()
	m.MockSprocGetVenue()
	m.MockSprocUpdateUsedCapacity(-(guest.AdditionalGuests + 1))
	err := DeleteGuest(guest)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteGuest_FailGetGuest(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName_NoGuest()
	err := DeleteGuest(guest)
	if err == nil {
		t.Fatal(err)
	}
}

func TestDeleteGuest_FailRemove(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocRemoveGuest_Fail()
	err := DeleteGuest(guest)
	if err == nil {
		t.Fatal(err)
	}
}

func TestGuestArrival(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable(guest.TableNumber)
	m.MockSprocGetVenue()
	m.MockSprocUpdateUsedCapacity(0)
	m.MockSprocGuestArrived()
	m.MockSprocGetGuestByName()
	_, err := GuestArrival(guest)
	if err != nil {
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

func TestGuestArrival_FailUpdateCapacity(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable(guest.TableNumber)
	m.MockSprocGetVenue()
	m.MockSprocUpdateUsedCapacity_Error(0)
	_, err := GuestArrival(guest)
	if err == nil {
		t.Fatal(err)
	}
}

func TestGuestLeaves(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGuestLeft()
	err := GuestLeaves(guest)
	if err != nil {
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

func TestGetSpaceOnTable(t *testing.T){
	m := database.NewMock()
	m.MockSprocGetGuestsAtTable(1)
	_, err := getSpaceOnTable(1, v)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetSpaceOnTable_Fail(t *testing.T){
	m := database.NewMock()
	m.MockSprocGetGuestsAtTable_NoGuests(1)
	_, err := getSpaceOnTable(1, v)
	if err == nil {
		t.Fatal(err)
	}
}

func TestFindTableFor(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestsAtTable(1)
	_, err := findTableFor(1, v)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFindTableFor_FailNoGuests(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestsAtTable_NoGuests(1)
	_, err := findTableFor(1, v)
	if err == nil {
		t.Fatal(err)
	}
}

func TestFindTableFor_FailAllTablesFull(t *testing.T) {
	m := database.NewMock()
	for i := 1; i <= v.NumberOfTables; i++ {
		m.MockSprocGetGuestsAtTable_FullTable(i)
	}
	m.MockSprocGetGuestsAtTable_FullTable(1)
	_, err := findTableFor(1, v)
	if err == nil {
		t.Fatal(err)
	}
}