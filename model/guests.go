package model

import (
	"errors"
	"github.com/benacook/getGround-technical-task/model/data"
	"github.com/benacook/getGround-technical-task/model/database"
	"log"
	"math"
)


//======================================================================================

//getSpaceOnTable returns the space on the given table.
//returns an error if it fails to get the guests on the table from the database.
func getSpaceOnTable(table int, venue data.Venue) (int, error) {
	guests, err := database.SprocGetGuestsAtTable(table)
	if err != nil {
		log.Println(err)
		err = errors.New("error: failed to get guests at the table")
		return 0, err
	}

	guestsAtTable := 0
	for _, guest := range guests {
		guestsAtTable += guest.AdditionalGuests + 1
	}

	return venue.TableSize - guestsAtTable, err

}

//======================================================================================
//findTableFor finds a table that can accommodate the given number of guests at the
//given venue.
//returns the table number of the table the guest and their additional guests will
//fit on. If they wont fit on any table, an error is returned.
func findTableFor(numGuests int, venue data.Venue) (int, error){
	for i := 1; i <= venue.NumberOfTables; i++ {
		space, err := getSpaceOnTable(i, venue)
		if err != nil {
			return 0, err
		}

		if space - numGuests >= 0{
			return i, nil
		}
	}
	return 0, errors.New("no table avaliable for number of guests")
}

//======================================================================================

//AddGuest adds a guest to the guest list,
//as long as there is space in the venue for them and their additional guests,
//and there is a table that they can all fit on together.
//It will attempt to put the guests on the table specified in the Guest struct,
//but if this is not possible it will try to fit them on a different table.
//returns a guest struct with the assigned table number,
//which may different from the requested.
//Wil return an error if it cant get the venue details, the venue is full,
//there is no table to accommodate the guest and their additional guests,
//or adding the guest to the guest list in the database fails.
func AddGuest(guest data.Guest) (data.Guest, error) {

	v, err := GetVenue()
	if err != nil {
		log.Println(err)
		err = errors.New("error: could not get venue, maybe no venue exists")
		return data.Guest{}, err
	}

	//the venue is full
	if v.UsedCapacity == v.Capacity{
		return data.Guest{}, errors.New("venue full")
	}

	space, err := getSpaceOnTable(guest.TableNumber, v)
	if err != nil {
		return data.Guest{}, err
	}

	if space - (guest.AdditionalGuests + 1) < 0{
		guest.TableNumber, err = findTableFor(guest.AdditionalGuests + 1, v)
		if err != nil {
			log.Println(err)
			return data.Guest{}, err
		}
	}

	err = database.SprocAddGuest(guest)
	if err != nil {
		log.Println(err)
		err = errors.New("error: failed to add guest")
		return data.Guest{}, err
	}

	g, err := GetGuest(guest.Name)

	VenueAddToUsedCapacity(g.AdditionalGuests + 1)

	return g, nil
}

//======================================================================================
//GetGuest takes the name of a guest and gets their details from the database.
//Returns the requests guest if successful.
//Returns an error if the guest does not exist,
//or if there is an issue contacting the database.
func GetGuest(name string) (data.Guest, error) {
	g, err := database.SprocGetGuestByName(name)
	if err != nil{
		log.Println(err)
		err = errors.New("error: no guest by that name")
		return data.Guest{}, err
	}
	return g, nil
}

//======================================================================================
//GetAllGuests returns all of the guests on the guest list.
//Returns an empty slice of Guest if the list is empty.
//Returns an error if there is an issue getting the list.
func GetAllGuests() ([]data.Guest, error){
	guests, err := database.SprocGetAllGuests()
	if err != nil {
		log.Println(err)
		err = errors.New("error: could not get guest list")
		return []data.Guest{}, err
	}
	return guests, nil
}

//======================================================================================
//GetArrivedGuests returns all of the guests on the guest list marked as arrived.
//Returns an empty slice of Guest if no-one has arrived.
//Returns an error if there is an issue getting the list.
func GetArrivedGuests() ([]data.Guest, error){
	guests, err := database.SprocGetArrivedGuests()
	if err != nil {
		log.Println(err)
		err = errors.New("error: could not get arrived guests")
		return []data.Guest{}, err
	}
	return guests, nil
}

//======================================================================================
//DeleteGuest removes a guest from the guest list in the database.
//First it makes sure that the guest exists and will error if they dont.
//Returns an error if it failed to delete the guest or if they dont exist in the
//database before deletion.
func DeleteGuest(guest data.Guest) error {
	g, err:= GetGuest(guest.Name)
	if err != nil {
		log.Println(err)
		return err
	}
	err = database.SprocRemoveGuest(g.Name)
	if err != nil {
		err = errors.New("error: could not remove guest")
		log.Println(err)
		return err
	}
	VenueSubtractFromUsedCapacity(g.AdditionalGuests + 1)

	return nil
}

//======================================================================================
//GuestArrival marks a guest as arrived at the party.
//If the number of additional guests is greater than expected,
//it will accept the guest if their table can accommodate their actual number of
//additional guests. Otherwise it will return an error.
//Returns the guest with their new number of additional guests, arrival time,
//and confirms they have been marked as arrived.
//returns an error if the guest is not on the list, it fails to get the venue,
//the guest and additional guests cannot fit on the table,
//or marking the guest as arrived in the database fails.
func GuestArrival(guest data.Guest) (data.Guest, error) {
	g, err := GetGuest(guest.Name)
	if err != nil {
		return data.Guest{}, err
	}

	v, err := GetVenue()
	if err != nil {
		return data.Guest{}, err
	}

	space, err := getSpaceOnTable(g.TableNumber, v)
	extraGuests := int(math.Abs(float64(g.AdditionalGuests - guest.
		AdditionalGuests)))

	if extraGuests > space{
		return data.Guest{}, errors.New(
			"error: not all guests can fit on the table")
	}

	if err := VenueAddToUsedCapacity(extraGuests); err != nil{
		return data.Guest{}, err
	}

	err = database.SprocGuestArrived(guest.Name, guest.AdditionalGuests)
	if err != nil {
		log.Println(err)
		err = errors.New("error: could not mark guest as arrived")
		return data.Guest{}, err
	}

	g, err = GetGuest(guest.Name)
	if err != nil {
		return data.Guest{}, err
	}
	return g, nil
}

//======================================================================================
//GuestLeaves un-marks the guest as arrived in the database,
//but does not modify the arrival time.
//Returns an error if the operation fails or if the guest was not invited.
func GuestLeaves(guest data.Guest) error {
	_, err := database.SprocGetGuestByName(guest.Name)
	if err != nil {
		log.Println(err)
		err = errors.New("error: could not get guest by that name")
		return err
	}
	err = database.SprocGuestLeft(guest.Name)
	if err != nil {
		err = errors.New("error: failed to mark guest as departed")
		log.Println(err)
		return err
	}
	return nil
}
