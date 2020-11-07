package database

import (
	"github.com/benacook/GetGround-Assignment/model/data"
	"log"
)

//======================================================================================

//SprocAddVenue calls the stored procedure AddVenue to get add the venue to the
//database.
//This calculates the table size from the number of tables and venue capacity.
//If the capacity is not divisible by the number of tables (
//to get a whole number) then this will add more tables to achieve that.
//Returns an error if the call to th stored procedure was unsuccessful.
func SprocAddVenue(venue data.Venue) error{

	//add an extra tables if the capacity can not be
	//divided by the number of tables with a remainder of 0
	//you cant have 7.1 seats at a table!
	venue.TableSize = venue.Capacity / venue.NumberOfTables
	for (venue.Capacity % venue.NumberOfTables) != 0 {
		venue.TableSize = venue.Capacity / venue.NumberOfTables
		venue.NumberOfTables++
	}


	_, err := Db.Exec("call AddVenue(?, ?, ?, ?)",
		venue.Name,	venue.Capacity, venue.NumberOfTables, venue.TableSize)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//======================================================================================

//SprocGetVenue calls the stored procedure GetVenue to get the
// venue from the database.
//the returned row is then put into a Venue struct and returned.
//Returns an error if the call to th stored procedure was unsuccessful.
func SprocGetVenue() (data.Venue, error){
	rows, err := Db.Query("call GetVenue()")
	if err != nil {
		log.Println(err)
		return data.Venue{}, err
	}
	defer rows.Close()

	v := data.Venue{}
	rows.Next()

	if err := rows.Scan(&v.Id, &v.Name, &v.Capacity,
		&v.NumberOfTables, &v.TableSize, &v.NextFreeTable,
		&v.UsedCapacity); err != nil{
		return data.Venue{}, err
	}

	return v, nil
}

//======================================================================================

//SprocDeleteVenue calls the stored procedure DeleteVenue to remove the venue from
//the database.
//Returns an error if the call to th stored procedure was unsuccessful.
func SprocDeleteVenue() error {
	_, err := Db.Exec("call DeleteVenue()")
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//======================================================================================

//SprocVenueUpdateUsedCapacity calls the stored procedure VenueUpdateUsedCapacity to
//update the used capacity of the venue - i.e. assigned capacity from the guest list.
//Returns an error if the call to th stored procedure was unsuccessful.
func SprocVenueUpdateUsedCapacity(venueName string, newUsedCapacity int, ) error {
	_, err := Db.Exec("call UpdateUsedCapacity(?, ?)",
		venueName, newUsedCapacity)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
