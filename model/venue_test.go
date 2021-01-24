package model

import (
	"github.com/benacook/PartyPlanner-Assignment/model/data"
	"github.com/benacook/PartyPlanner-Assignment/model/database"
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

func TestAddVenue(t *testing.T) {
	m := database.NewMock()
	m.MockSprocAddVenue()
	m.MockSprocGetAllGuests()
	m.MockSprocUpdateUsedCapacity(200)
	m.MockSprocGetVenue()
	vv := v
	vv.UsedCapacity = 195

	_, err := AddVenue(vv)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestAddVenue_FailBadVenue(t *testing.T) {
	vv := v
	vv.Capacity = 0
	_, err := AddVenue(vv)
	if err == nil {
		t.Fatalf("expected bad venue, got: %v", err)
	}
}

func TestAddVenue_FailNoVenue(t *testing.T) {
	m := database.NewMock()
	m.MockSprocAddVenue_NoVenue()
	_, err := AddVenue(v)
	if err == nil {
		t.Fatalf("expected no venue, got: %v", err)
	}
}

func TestAddVenue_FailNoGuests(t *testing.T) {
	m := database.NewMock()
	m.MockSprocAddVenue()
	m.MockSprocGetAllGuests_NoGuests()
	_, err := AddVenue(v)
	if err == nil {
		t.Fatalf("expected no guests, got: %v", err)
	}
}

func TestAddVenue_FailUpdateCapacity(t *testing.T) {
	m := database.NewMock()
	m.MockSprocAddVenue()
	m.MockSprocGetAllGuests()
	m.MockSprocUpdateUsedCapacity_Error(5)
	_, err := AddVenue(v)
	if err == nil {
		t.Fatalf("expected fail update capacity, got: %v", err)
	}
}

func TestAddVenue_FailGetVenue(t *testing.T) {
	m := database.NewMock()
	m.MockSprocAddVenue()
	m.MockSprocGetAllGuests()
	m.MockSprocUpdateUsedCapacity(200)
	m.MockSprocGetVenue_NoVenue()
	vv := v
	vv.UsedCapacity = 195
	_, err := AddVenue(vv)
	if err == nil {
		t.Fatalf("expected fail get venue, got: %v", err)
	}
}

func TestVenueAddToUsedCapacity(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocUpdateUsedCapacity(5)
	err := VenueAddToUsedCapacity(5)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestVenueAddToUsedCapacity_FailNoVenue(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue_NoVenue()
	err := VenueAddToUsedCapacity(5)
	if err == nil {
		t.Fatalf("expected no venue, got: %v", err)
	}
}

func TestVenueAddToUsedCapacity_FailToUpdate(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocUpdateUsedCapacity_Error(5)
	err := VenueAddToUsedCapacity(5)
	if err == nil {
		t.Fatalf("expected error on update capacity, got: %v", err)
	}
}

func TestVenueSubtractFromUsedCapacity(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocUpdateUsedCapacity(0)
	err := VenueSubtractFromUsedCapacity(0)
	if err != nil {
		t.Fatalf("expected nil, got: %v", err)
	}
}

func TestVenueSubtractFromUsedCapacity_FailNoVenue(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue_NoVenue()
	err := VenueSubtractFromUsedCapacity(0)
	if err == nil {
		t.Fatalf("expected error on get venue, got: %v", err)
	}
}

func TestVenueSubtractFromUsedCapacity_FailToUpdate(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocUpdateUsedCapacity_Error(-5)
	err := VenueSubtractFromUsedCapacity(5)
	if err == nil {
		t.Fatalf("expected error on update capacity, got: %v", err)
	}
}

func TestGetRemainingSeats(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocGetArrivedGuests()
	_, err := GetRemainingSeats()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestGetRemainingSeats_FailNoGuests(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	m.MockSprocGetArrivedGuests_NoGuests()
	_, err := GetRemainingSeats()
	if err == nil {
		t.Fatalf("expected error on get arrived guests, got: %v", err)
	}
}

func TestGetRemainingSeats_FailNoVenue(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue_NoVenue()
	_, err := GetRemainingSeats()
	if err == nil {
		t.Fatalf("expected error on get venue, got: %v", err)
	}
}

func TestGetRemainingBookableSeats(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue()
	_, err := GetRemainingBookableSeats()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestGetRemainingBookableSeats_Fail(t *testing.T) {
	m := database.NewMock()
	m.MockSprocGetVenue_NoVenue()
	_, err := GetRemainingBookableSeats()
	if err == nil {
		t.Fatalf("expected error on get venue, got: %v", err)
	}
}

func TestDeleteVenue(t *testing.T) {
	m := database.NewMock()
	m.MockSprocDeleteVenue()
	err := DeleteVenue()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestDeleteVenue_Fail(t *testing.T) {
	m := database.NewMock()
	m.MockSprocDeleteVenue_NoVenue()
	err := DeleteVenue()
	if err == nil {
		t.Fatalf("expected error on get venue, got: %v", err)
	}
}