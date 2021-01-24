package model

import (
	"github.com/benacook/PartyPlanner-Assignment/model/data"
	"github.com/benacook/PartyPlanner-Assignment/model/database"
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
	g, err := AddGuest(guest)
	if err != nil {
		t.Fatal(err)
	}
	if g.Name != guest.Name{
		t.Fatal("returned guest name did not match the guest added")
	}
	if g.AdditionalGuests != guest.AdditionalGuests {
		t.Fatal("returned guest's additional guests did not match the guest added")
	}
}

func TestAddGuest_FailAdd(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable(guest.TableNumber)
	m.MockSprocAddGuest_Error()
	g, err := AddGuest(guest)
	if err == nil {
		t.Fatal(err)
	}
	blank := data.Guest{}
	if g != blank {
		t.Fatal("expected blank guest")
	}
}

func TestAddGuest_FailGetAddedGuest(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable(guest.TableNumber)
	m.MockSprocAddGuest()
	m.MockSprocGetGuestByName_NoGuest()
	m.MockSprocUpdateUsedCapacity(guest.AdditionalGuests + 1)
	g, err := AddGuest(guest)
	if err == nil {
		t.Fatal(err)
	}
	blank := data.Guest{}
	if g != blank {
		t.Fatal("expected blank guest")
	}
}

func TestAddGuest_FailNoVenue(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue_NoVenue()
	g, err := AddGuest(guest)
	if err == nil {
		t.Fatal(err)
	}
	blank := data.Guest{}
	if g != blank {
		t.Fatal("expected blank guest")
	}
}

func TestAddGuest_FailFullVenue(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue_FullVenue()
	g, err := AddGuest(guest)
	if err == nil {
		t.Fatal(err)
	}
	blank := data.Guest{}
	if g != blank {
		t.Fatal("expected blank guest")
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
	g, err := AddGuest(g)
	if err == nil {
		t.Fatal(err)
	}
	blank := data.Guest{}
	if g != blank {
		t.Fatal("expected blank guest")
	}
}

func TestAddGuest_FailNoSpace(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable_NoGuests(guest.TableNumber)
	for i := 1; i <= v.NumberOfTables; i++{
		m.MockSprocGetGuestsAtTable_NoGuests(guest.TableNumber)
	}
	g, err := AddGuest(guest)
	if err == nil {
		t.Fatal(err)
	}
	blank := data.Guest{}
	if g != blank {
		t.Fatal("expected blank guest")
	}
}

func TestGetArrivedGuests(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetArrivedGuests()
	g, err := GetArrivedGuests()

	if err != nil {
		t.Fatal(err)
	}
	gs := []data.Guest{}
	gs = append(gs, guest)

	if g[0].Name != gs[0].Name {
		t.Fatal("guest names dont match")
	}
	if g[0].AdditionalGuests != gs[0].AdditionalGuests {
		t.Fatal("additional guests dont match")
	}

}

func TestGetArrivedGuests_Fail(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetArrivedGuests_NoGuests()
	g, err := GetArrivedGuests()

	if err == nil {
		t.Fatal(err)
	}

	if len(g) > 0{
		t.Fatal("expected empty guests")
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
	g, err := GuestArrival(guest)
	if err != nil {
		t.Fatal(err)
	}
	if g.Name != guest.Name{
		t.Fatal("returned guest name did not match the guest added")
	}
	if g.AdditionalGuests != guest.AdditionalGuests {
		t.Fatal("returned guest's additional guests did not match the guest added")
	}

}

func TestGuestArrival_FailTable(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable_NoGuests(guest.TableNumber)
	g, err := GuestArrival(guest)
	if err == nil {
		t.Fatal(err)
	}

	blank := data.Guest{}
	if g != blank {
		t.Fatal("expected blank guest")
	}
}

func TestGuestArrival_FailNoGuest(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName_NoGuest()
	g, err := GuestArrival(guest)
	if err == nil {
		t.Fatal(err)
	}

	blank := data.Guest{}
	if g != blank {
		t.Fatal("expected blank guest")
	}
}

func TestGuestArrival_FailNoVenue(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGetVenue_NoVenue()
	g, err := GuestArrival(guest)

	if err == nil {
		t.Fatal(err)
	}

	blank := data.Guest{}
	if g != blank {
		t.Fatal("expected blank guest")
	}
}

func TestGuestArrival_FailNoSpace(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable(guest.TableNumber)
	g := guest
	g.AdditionalGuests = 100

	g, err := GuestArrival(g)
	if err == nil {
		t.Fatal(err)
	}

	blank := data.Guest{}
	if g != blank {
		t.Fatal("expected blank guest")
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

	g, err := GuestArrival(guest)
	if err == nil {
		t.Fatal(err)
	}

	blank := data.Guest{}
	if g != blank {
		t.Fatal("expected blank guest")
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

	g, err := GuestArrival(guest)
	if err == nil {
		t.Fatal(err)
	}

	blank := data.Guest{}
	if g != blank {
		t.Fatal("expected blank guest")
	}
}

func TestGuestArrival_FailUpdateCapacity(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetGuestByName()
	m.MockSprocGetVenue()
	m.MockSprocGetGuestsAtTable(guest.TableNumber)
	m.MockSprocGetVenue()
	m.MockSprocUpdateUsedCapacity_Error(0)

	g, err := GuestArrival(guest)
	if err == nil {
		t.Fatal(err)
	}

	blank := data.Guest{}
	if g != blank {
		t.Fatal("expected blank guest")
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