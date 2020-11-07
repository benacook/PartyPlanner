package database

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/benacook/GetGround-Assignment/model/data"
	_"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type DbMock struct {
	db   *sql.DB
	mock sqlmock.Sqlmock
}

var (
	g = data.Guest{
		Id: 1,
		Name: "Ben",
		AdditionalGuests: 4,
		TableNumber: 1,
		Arrived: false,
	}

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

//======================================================================================

//NewMock returns a DbMock object for mocking th database connection.
//a method is provided for mocking all of the stored procedures for both success
//states and error states.
//Success mocks are named .MockSproc<sproc name>.
//Error state mockas are named .MockSproc<sproc name>_<error description>
func NewMock() DbMock {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}
	Db.Db = db
	mockDb := DbMock{db, mock}
	return mockDb
}

//======================================================================================

//mockGuestRows returns a mocked sql row for a guest that has not arrived at the party.
func mockGuestRows() *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"id", "name",
		"additionalguests", "tablenumber", "arrived", "arrivaltime"}).
		AddRow(g.Id, g.Name, g.AdditionalGuests, g.TableNumber, g.Arrived, g.ArrivalTime)

	return rows
}

//======================================================================================

//mockGuestRowsArrived returns a mocked sql row for a guest that has arrived at the party.
func mockGuestRowsArrived() *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"id", "name",
		"additionalguests", "tablenumber", "arrived", "arrivaltime"}).
		AddRow(g.Id, g.Name, g.AdditionalGuests, g.TableNumber, true, time.Now().String())

	return rows
}

//======================================================================================

//mockVenueRows returns a mocked sql row for the venue.
func mockVenueRows() *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"id", "name",
		"capacity", "numberoftables", "tablesize", "nextfreetable", "usedcapacity"}).
		AddRow(v.Id, v.Name, v.Capacity, v.NumberOfTables, v.TableSize,
			v.NextFreeTable, v.UsedCapacity)
	return rows
}

//======================================================================================
func (dbMock *DbMock) MockSprocGetGuestByName()  {
	query := "call GetGuestByName(?)"
	rows := mockGuestRows()
	dbMock.mock.ExpectQuery(query).WithArgs(g.Name).WillReturnRows(rows).
		WillReturnError(nil)
}

//======================================================================================
func (dbMock *DbMock) MockSprocGetGuestByName_NoGuest()  {
	query := "call GetGuestByName(?)"
	dbMock.mock.ExpectQuery(query).WithArgs(g.Name).WillReturnError(sql.ErrNoRows).
		WillReturnRows(sqlmock.NewRows([]string{""}))
}

//======================================================================================
func (dbMock *DbMock) MockSprocGetVenue() {
	query := "call GetVenue()"
	rows := mockVenueRows()
	dbMock.mock.ExpectQuery(query).WillReturnRows(rows).WillReturnError(nil)
}

//======================================================================================
func (dbMock *DbMock) MockSprocGetVenue_NoVenue() {
	query := "call GetVenue()"
	dbMock.mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)
}

//======================================================================================
func (dbMock *DbMock) MockSprocGetGuestsAtTable(table int) {
	query := "call GetGuestsAtTable(?)"
	rows := mockGuestRows()
	dbMock.mock.ExpectQuery(query).WithArgs(table).WillReturnRows(rows).
		WillReturnError(nil)
}

//======================================================================================
func (dbMock *DbMock) MockSprocGetGuestsAtTable_NoGuests(table int) {
	query := "call GetGuestsAtTable(?)"
	dbMock.mock.ExpectQuery(query).WithArgs(table).WillReturnError(sql.ErrNoRows)
}

//======================================================================================
func (dbMock *DbMock) MockSprocGuestArrived() {
	query := "call GuestArrived(?, ?)"
	dbMock.mock.ExpectExec(query).WithArgs(g.Name, g.AdditionalGuests).
		WillReturnError(nil).WillReturnResult(
		sqlmock.NewResult(1, 1))
}

//======================================================================================
func (dbMock *DbMock) MockSprocGuestArrived_NoGuest() {
	query := "call GuestArrived(?, ?)"
	dbMock.mock.ExpectExec(query).WithArgs(g.Name, g.AdditionalGuests).
		WillReturnError(sql.ErrNoRows).WillReturnResult(
		sqlmock.NewErrorResult(errors.New("guest not on list")))
}

//======================================================================================
func (dbMock *DbMock) MockSprocGuestLeft() {
	query := "call GuestLeft(?)"
	dbMock.mock.ExpectExec(query).WithArgs(g.Name).
		WillReturnError(nil).WillReturnResult(
		sqlmock.NewResult(1, 1))
}

//======================================================================================
func (dbMock *DbMock) MockSprocGuestLeft_NoGuest() {
	query := "call GuestLeft(?)"
	dbMock.mock.ExpectExec(query).WithArgs(g.Name).
		WillReturnError(sql.ErrNoRows).WillReturnResult(
		sqlmock.NewErrorResult(errors.New("guest was not invited")))
}

//======================================================================================
func (dbMock *DbMock) MockSprocGetArrivedGuests() {
	query := "call GetArrivedGuests()"
	rows := mockGuestRowsArrived()
	dbMock.mock.ExpectQuery(query).WillReturnRows(rows).WillReturnError(nil)
}

//======================================================================================
func (dbMock *DbMock) MockSprocGetArrivedGuests_NoGuests() {
	query := "call GetArrivedGuests()"
	dbMock.mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{""})).
		WillReturnError(sql.ErrNoRows)
}

//======================================================================================
func (dbMock *DbMock) MockSprocAddGuest() {
	query := "call AddGuest(?, ?, ?)"

	dbMock.mock.ExpectExec(query).WithArgs(g.Name, g.AdditionalGuests,
		g.TableNumber).WillReturnError(nil).WillReturnResult(
		sqlmock.NewResult(1, 1))
}

//======================================================================================
func (dbMock *DbMock) MockSprocAddGuest_Error() {
	query := "call AddGuest(?, ?, ?)"

	dbMock.mock.ExpectExec(query).WithArgs(g.Name, g.AdditionalGuests,
		g.TableNumber).WillReturnError(sql.ErrNoRows).WillReturnResult(
		sqlmock.NewErrorResult(errors.New("failed to update capaicty")))

}

//======================================================================================
func (dbMock *DbMock) MockSprocUpdateUsedCapacity(capacity int) {
	query := "call UpdateUsedCapacity(?, ?)"

	dbMock.mock.ExpectExec(query).WithArgs(v.Name, capacity).WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

//======================================================================================
func (dbMock *DbMock) MockSprocUpdateUsedCapacity_Error(capacity int) {
	query := "call UpdateUsedCapacity(?, ?)"

	dbMock.mock.ExpectExec(query).WithArgs(v.Name, capacity).WillReturnError(sql.ErrNoRows).
		WillReturnResult(sqlmock.NewErrorResult(errors.New(
			"failed to update capaicty")))
}

//======================================================================================
func (dbMock *DbMock) MockSprocRemoveGuest() {
	query := "call RemoveGuest(?)"
	dbMock.mock.ExpectExec(query).WithArgs(g.Name).WillReturnError(nil).WillReturnResult(
		sqlmock.NewResult(1, 1))
}

//======================================================================================
func (dbMock *DbMock) MockSprocRemoveGuest_NoGuest() {
	query := "call RemoveGuest(?)"
	dbMock.mock.ExpectExec(query).WithArgs(g.Name).WillReturnError(sql.ErrNoRows).
		WillReturnResult(
		sqlmock.NewErrorResult(errors.New("no guest by that name")))
}

//======================================================================================
func (dbMock *DbMock) MockSprocGetAllGuests() {
	query := "call GetAllGuests()"
	rows := mockGuestRows()
	dbMock.mock.ExpectQuery(query).WillReturnRows(rows).WillReturnError(nil)
}

//======================================================================================
func (dbMock *DbMock) MockSprocGetAllGuests_NoGuests() {
	query := "call GetAllGuests()"
	dbMock.mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows).
		WillReturnRows(sqlmock.NewRows([]string{""}))
}

//======================================================================================
func (dbMock *DbMock) MockSprocAddVenue() {
	query := "call AddVenue(?, ?, ?, ?)"

	dbMock.mock.ExpectExec(query).WithArgs(v.Name, v.Capacity, v.NumberOfTables,
		v.TableSize).WillReturnError(nil).WillReturnResult(
		sqlmock.NewResult(1, 1))
}

//======================================================================================
func (dbMock *DbMock) MockSprocAddVenue_NoVenue() {
	query := "call AddVenue(?, ?, ?, ?)"

	dbMock.mock.ExpectExec(query).WithArgs(v.Name, v.Capacity, v.NumberOfTables,
		v.TableSize).WillReturnError(sql.ErrNoRows).WillReturnResult(
		sqlmock.NewErrorResult(errors.New("no guest by that name")))
}

//======================================================================================
func (dbMock *DbMock) MockSprocDeleteVenue() {
	query := "call DeleteVenue()"

	dbMock.mock.ExpectExec(query).WillReturnError(nil).WillReturnResult(
		sqlmock.NewResult(1, 1))
}

//======================================================================================
func (dbMock *DbMock) MockSprocDeleteVenue_NoVenue() {
	query := "call DeleteVenue()"

	dbMock.mock.ExpectExec(query).WillReturnError(sql.ErrNoRows).WillReturnResult(
		sqlmock.NewErrorResult(errors.New("no venue to delete")))
}





