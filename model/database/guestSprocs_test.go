package database

import (
	"database/sql"
	"testing"
)

func TestSprocAddGuest(t *testing.T) {
	m := NewMock()
	m.MockSprocAddGuest()
	
	err := SprocAddGuest(g)
	if err != nil {
		t.Fatal(err)
	}
}


func TestSprocAddGuest_Fail(t *testing.T) {
	m := NewMock()
	m.MockSprocAddGuest_Error()

	err := SprocAddGuest(g)
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}
}

func TestSprocGetGuestByName(t *testing.T) {
	m := NewMock()
	m.MockSprocGetGuestByName()
	
	_, err := SprocGetGuestByName(g.Name)
	if err != nil {
		t.Fatal(err)
	}
}


func TestSprocGetGuestByName_Fail(t *testing.T) {
	m := NewMock()
	m.MockSprocGetGuestByName_NoGuest()

	_, err := SprocGetGuestByName("Ben")
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}
}

func TestSprocGetAllGuests(t *testing.T)  {
	m := NewMock()
	m.MockSprocGetAllGuests()

	r, err := SprocGetAllGuests()
	if err != nil {
		t.Fatal(err)
	}

	if r[0] != g{
		t.Fatal("incorrect data returned")
	}
}


func TestSprocGetAllGuests_Fail(t *testing.T)  {
	m := NewMock()
	m.MockSprocGetAllGuests_NoGuests()

	_, err := SprocGetAllGuests()
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}
}

func TestSprocGetArrivedGuests(t *testing.T) {
	m := NewMock()
	m.MockSprocGetArrivedGuests()

	_, err := SprocGetArrivedGuests()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSprocGetArrivedGuests_Fail(t *testing.T) {
	 m := NewMock()
	 m.MockSprocGetArrivedGuests_NoGuests()

	_, err := SprocGetArrivedGuests()
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}
}

func TestSprocRemoveGuest(t *testing.T) {
	 m := NewMock()
	 m.MockSprocRemoveGuest()

	err := SprocRemoveGuest(g.Name)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSprocRemoveGuest_Fail(t *testing.T) {
	 m := NewMock()
	 m.MockSprocRemoveGuest_NoGuest()

	err := SprocRemoveGuest(g.Name)
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}
}

func TestSprocGuestArrived(t *testing.T) {
	 m := NewMock()
	 m.MockSprocGuestArrived()

	err := SprocGuestArrived(g.Name, g.AdditionalGuests)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSprocGuestArrived_Fail(t *testing.T) {
	 m := NewMock()
	 m.MockSprocGuestArrived_NoGuest()

	err := SprocGuestArrived(g.Name, g.AdditionalGuests)
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}
}

func TestSprocGuestLeft(t *testing.T) {
	 m := NewMock()
	 m.MockSprocGuestLeft()

	err := SprocGuestLeft(g.Name)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSprocGuestLeft_Fail(t *testing.T) {
	 m := NewMock()
	 m.MockSprocGuestLeft_NoGuest()

	err := SprocGuestLeft(g.Name)
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}
}

func TestSprocGetGuestsAtTable(t *testing.T) {
	m := NewMock()
 	m.MockSprocGetGuestsAtTable()

	r, err := SprocGetGuestsAtTable(1)
	if err != nil {
		t.Fatal(err)
	}

	if r[0] != g{
		t.Fatal("incorrect data returned")
	}
}

func TestSprocGetGuestsAtTable_Fail(t *testing.T) {
	 m := NewMock()
	 m.MockSprocGetGuestsAtTable_NoGuests()

	_, err := SprocGetGuestsAtTable(5)
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}
}
