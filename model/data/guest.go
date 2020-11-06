package data

//======================================================================================

//Guest struct contains all of the data of a guest on the guest list,
//contains the same data as the guest_list table in the mysql database.
type Guest struct {
	Id int
	Name string
	AdditionalGuests int `json:"accompanying_guests"`
	TableNumber int
	Arrived bool
	ArrivalTime string
}

//======================================================================================

//GuestList contains a slice of Guest.
type GuestList struct {
	Guests []Guest
}