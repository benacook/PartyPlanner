package database

import (
	"database/sql"
	"testing"
)

func TestSprocAddVenue(t *testing.T) {
	vv := v
	vv.Capacity += 10

	m := NewMock()
	m.MockSprocAddVenue_WithCapacityTables(vv.Capacity, 21)
	err := SprocAddVenue(vv)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSprocAddVenue2(t *testing.T) {
	m := NewMock()
	m.MockSprocAddVenue()
	err := SprocAddVenue(v)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSprocAddVenue_Fail(t *testing.T) {
	m := NewMock()
	m.MockSprocAddVenue_NoVenue()

	err := SprocAddVenue(v)
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}
}

func TestSprocGetVenue(t *testing.T) {
	m := NewMock()
	m.MockSprocGetVenue()

	_, err := SprocGetVenue()
	if err != nil {
		t.Fatal(err)
	}
}


func TestSprocGetVenue_Fail(t *testing.T) {
	m := NewMock()
	m.MockSprocGetVenue_NoVenue()

	_, err := SprocGetVenue()
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}
}


func TestSprocDeleteVenue(t *testing.T) {
	m := NewMock()
	m.MockSprocDeleteVenue()

	err := SprocDeleteVenue()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSprocDeleteVenue_Fail(t *testing.T) {
	m := NewMock()
	m.MockSprocDeleteVenue_NoVenue()

	err := SprocDeleteVenue()
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}
}

func TestSprocVenueUpdateUsedCapacity(t *testing.T) {
	m := NewMock()
	m.MockSprocUpdateUsedCapacity(v.Capacity)

	err := SprocVenueUpdateUsedCapacity(v.Name, v.Capacity)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSprocVenueUpdateUsedCapacity_Fail(t *testing.T) {
	m := NewMock()
	m.MockSprocUpdateUsedCapacity_Error(v.Capacity)

	err := SprocVenueUpdateUsedCapacity(v.Name, v.Capacity)
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}
}

