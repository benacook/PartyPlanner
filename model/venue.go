package model

import (
	"errors"
	"github.com/benacook/GetGround-Assignment/model/data"
	"github.com/benacook/GetGround-Assignment/model/database"
	"log"
)

//======================================================================================

//AddVenue adds the venue to the database.
//As a user can add guests to the guest list and then delete and add a different
//venue, this will update the used capacity of the venue immediately adding the venue
//to the database, it will not however check that all guests can fit on their assinged
//tables.
//Returns the venue's data (some if it is calculated in the call to the sproc).
//Returns an error if the venue is invalid, or failed to add the venue,
//or failed to calculate and update the capacity from the current guest list.
func AddVenue(venue data.Venue) (data.Venue, error)  {
	if venue.Name == "" ||
		venue.NumberOfTables < 1 ||
		venue.Capacity < 1 ||
		venue.Capacity < venue.NumberOfTables{
		return data.Venue{}, errors.New("invalid venue")
	}

	if err := database.SprocAddVenue(venue); err != nil{
		log.Println(err)
		err = errors.New(" could not add venue")
		return data.Venue{}, err
	}

	guests, err :=  GetAllGuests()
	if err != nil {
		log.Println(err)
		return data.Venue{}, err
	}

	for _, guest := range guests {
		venue.UsedCapacity += guest.AdditionalGuests + 1
	}

	err = database.SprocVenueUpdateUsedCapacity(venue.Name, venue.UsedCapacity)
	if err != nil{
		log.Println(err)
		err = errors.New(" could not update used capacity" +
			"from existing guest list")
		return data.Venue{}, err
	}

	v, err := GetVenue()
	if err != nil{
		return data.Venue{}, err
	}
	return v, nil
}

//======================================================================================

//GetVenue gets the venue from the database.
//Returns the venue.
//Returns an error if it failed to get the venue, or there is no venue.
func GetVenue() (data.Venue, error)  {
	v, err := database.SprocGetVenue()
	if err != nil{
		log.Println(err)
		err = errors.New(" could not get venue")
		return data.Venue{}, err
	}
	return v, nil
}

//======================================================================================

//VenueAddToUsedCapacity increases the venue's used capacity by the given amount.
//Returns an error if it fails to get the venue or fails to update the new used
//capacity in the database.
func VenueAddToUsedCapacity(guests int) error{
	v, err := GetVenue()
	if err != nil {
		return err
	}

	v.UsedCapacity += guests
	err = database.SprocVenueUpdateUsedCapacity(v.Name, v.UsedCapacity)
	if err != nil{
		log.Println(err)
		err = errors.New(" could not update used venue capacity")
		return err
	}
	return nil
}

//======================================================================================

//VenueSubtractFromUsedCapacity decreases the venue's used capacity by the given amount.
//Returns an error if it fails to get the venue or fails to update the new used
//capacity in the database.
func VenueSubtractFromUsedCapacity(guests int) error{
	v, err := GetVenue()
	if err != nil {
		return err
	}

	v.UsedCapacity -= guests
	err = database.SprocVenueUpdateUsedCapacity(v.Name, v.UsedCapacity)
	if err != nil{
		log.Println(err)
		err = errors.New(" could not update used venue capacity")
		return err
	}
	return nil
}

//======================================================================================

//DeleteVenue removed the venue from the database.
//Returns an error if the operation fails.
func DeleteVenue() error  {
	err := database.SprocDeleteVenue()
	if err != nil{
		log.Println(err)
		err = errors.New(" could not delete venue")
		return err
	}
	return nil
}

//======================================================================================

//GetRemainingSeats returns the number of empty seats during the party.
//Returns an error if it cannot get the venue or arrived guests.
func GetRemainingSeats() (int, error) {
	v, err := GetVenue()
	if err != nil {
		return 0, err
	}

	guests, err := GetArrivedGuests()
	if err != nil {
		return 0, err
	}

	total := 0
	for _, guest := range guests {
		total += guest.AdditionalGuests + 1
	}

	return v.Capacity - total, nil
}

//======================================================================================

//GetRemainingBookableSeats returns the number of remaining seats at the venue that
//guest can be invited to fill, i.e. remaining bookable capacity.
//Returns an error if it cannot get the venue.
func GetRemainingBookableSeats() (int, error) {
	v, err := GetVenue()
	if err != nil {
		return 0, err
	}
	return v.Capacity - v.UsedCapacity, nil
}
