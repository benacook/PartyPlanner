package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


const (
	dbUser string = "party-app"
	dbPassword string = "4321"
	dbName string = "getground-party"
	dbConStr string  =
	dbUser + ":"+ dbPassword + "" +
		"@tcp(127.0.0.1:3306)/" + dbName
)

var (
	 Db database
)

type database struct {
	Db *sql.DB
}

//======================================================================================

//Init initialises the global variable Db for general database operations.
func Init() error {
	var err error
	Db, err = NewDatabase("mysql", dbConStr)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

//======================================================================================

//NewDatabase creates a new connection to a database.
//Returns a useable database struct and nil error if successful.
//Returns a blank database struct and error if failure occurs.
func NewDatabase(sqlType, connectionString string) (database, error) {
	var err error
	Db, err := sql.Open(sqlType, connectionString)
	if err != nil {
		log.Println(err)
		return database{}, err
	}

	err = Db.Ping()
	if err != nil {
		log.Println(err)
		return database{}, err
	}

	return database{Db}, nil
}

//======================================================================================

//Close closes the database connection.
func (db *database)  Close(){
	db.Db.Close()
}

//======================================================================================

//Exec wraps the sql.Db.Exec method.
func (db *database) Exec(query string, args... interface{}) (sql.Result, error) {
	r, err := db.Db.Exec(query, args...)
	if err != nil{
		log.Println(err)
		return nil, err
	}
	return r, nil
}

//======================================================================================

//Query wraps the sql.Db.Query method.
func (db *database) Query(query string, args... interface{}) (*sql.Rows, error) {
	rows, err := db.Db.Query(query, args...)
	if err != nil{
		log.Println(err)
		return &sql.Rows{}, err
	}
	return rows, nil
}