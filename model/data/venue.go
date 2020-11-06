package data

//======================================================================================

//Venue struct contains all of the data of the venue,
//contains the same data as the venue table in the mysql database.
type Venue struct {
	Id int
	Name string
	Capacity int
	NumberOfTables int
	TableSize int
	NextFreeTable int
	UsedCapacity int
}

//======================================================================================

//EmptySeats struct is used to return empty seats,
//either bookable or empty at the party.
type EmptySeats struct {
	SeatsEmpty int `json:"seats_empty"`
}
