package database

import (
	"errors"
	"github.com/benacook/GetGround-Assignment/model/data"
	"log"
)

//======================================================================================

//SprocAddGuest calls the stored procedure AddGuest to insert a new Guest into the
//database.
//Returns an error if the call to th stored procedure was unsuccessful.
func SprocAddGuest(guest data.Guest) error {
	_, err := Db.Exec("call AddGuest(?, ?, ?)",
		guest.Name,	guest.AdditionalGuests, guest.TableNumber)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//======================================================================================

//SprocGetGuestByName calls the stored procedure GetGuestByName to get a guest
//given their name from database.
//The returned row is then put into a Guest struct and is returned.
//Returns an error if the call to th stored procedure was unsuccessful.
func SprocGetGuestByName(name string) (data.Guest, error) {
	rows, err := Db.Query("call GetGuestByName(?)", name)
	if err != nil {
		log.Println(err)
		return data.Guest{}, err
	}
	defer rows.Close()

	g := data.Guest{}
	rows.Next()

	rows.Scan(&g.Id, &g.Name, &g.AdditionalGuests,
		&g.TableNumber, &g.Arrived, &g.ArrivalTime)

	if g.Name == "" {
		return data.Guest{}, errors.New("no guest by that name")
	}

	return g, nil
}

//======================================================================================

//SprocGetAllGuests calls the stored procedure GetAllGuests to get a list of guest
//from database.
//The returned rows are then put into a slice of Guest and is returned.
//Returns an error if the call to th stored procedure was unsuccessful.
func SprocGetAllGuests() ([]data.Guest, error) {
	rows, err := Db.Query("call GetAllGuests()")
	if err != nil {
		log.Println(err)
		return []data.Guest{}, err
	}
	defer rows.Close()

	guests := []data.Guest{}
	for rows.Next(){
		g := data.Guest{}
		rows.Scan(&g.Id, &g.Name, &g.AdditionalGuests,
			&g.TableNumber, &g.Arrived, &g.ArrivalTime)
		guests = append(guests, g)
	}
	return guests, nil
}

//======================================================================================

//SprocGetArrivedGuests calls the stored procedure GetArrivedGuests to get a list of guest
// that have arrived at the party from database.
//The returned rows are then put into a slice of Guest and is returned.
//Returns an error if the call to th stored procedure was unsuccessful.
func SprocGetArrivedGuests() ([]data.Guest, error) {
	rows, err := Db.Query("call GetArrivedGuests()")
	if err != nil {
		log.Println(err)
		return []data.Guest{}, err
	}
	defer rows.Close()

	guests := []data.Guest{}
	for rows.Next(){
		g := data.Guest{}
		rows.Scan(&g.Id, &g.Name, &g.AdditionalGuests,
			&g.TableNumber, &g.Arrived, &g.ArrivalTime)
		guests = append(guests, g)
	}
	return guests, nil

}

//======================================================================================

//SprocRemoveGuest calls the stored procedure RemoveGuest to remove a guest with the
//given name from the database.
//Returns an error if the call to th stored procedure was unsuccessful.
func SprocRemoveGuest(name string) error {
	_ , err := Db.Exec("call RemoveGuest(?)", name)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

//======================================================================================

//SprocGuestArrived calls the stored procedure GuestArrived to update the given guest
//in the database with the number of additional guests and mark them as arrived.
//Returns an error if the call to th stored procedure was unsuccessful.
func SprocGuestArrived(name string, guests int) error {
	_, err := Db.Exec("call GuestArrived(?, ?)", name, guests)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//======================================================================================

//SprocGuestLeft calls the stored procedure GuestLeft to update the given guest in the
//database and mark them as departed.
//Returns an error if the call to th stored procedure was unsuccessful.
func SprocGuestLeft(name string) error {
	_, err := Db.Exec("call GuestLeft(?)", name)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

//======================================================================================

//SprocGetGuestsAtTable calls the stored procedure GetGuestsAtTable to get the
// guests assigned to the given table.
//the returned rows are then put into a slice of Guest and returned.
//Returns an error if the call to th stored procedure was unsuccessful.

func SprocGetGuestsAtTable(table int) ([]data.Guest, error) {
	rows, err := Db.Query("call GetGuestsAtTable(?)", table)
	if err != nil {
		log.Println(err)
		return []data.Guest{}, err
	}
	defer rows.Close()

	guests := []data.Guest{}
	for rows.Next(){
		g := data.Guest{}
		rows.Scan(&g.Id, &g.Name, &g.AdditionalGuests,
			&g.TableNumber, &g.Arrived, &g.ArrivalTime)
		guests = append(guests, g)
	}
	return guests, nil
}