package database

import "testing"

//======================================================================================
func TestNewDatabase(t *testing.T) {
	db, err := NewDatabase("mysql", dbConStr)
	db.Close()
	if err != nil {
		t.Fatal(err)
	}
}

//======================================================================================
func TestNewDatabase_Fail(t *testing.T) {
	_, err := NewDatabase("mysql", "")
	if err == nil {
		t.Fatal(err)
	}
}

//======================================================================================
func TestNewDatabase_Fail2(t *testing.T) {
	_, err := NewDatabase("", dbConStr)
	if err == nil {
		t.Fatal(err)
	}
}

//======================================================================================
func TestDatabase_Close(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("The went into panic")
		}
	}()
	db, err := NewDatabase("mysql", dbConStr)
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
}

//======================================================================================
//TestDatabase_Close_Fail causes panic.
func TestDatabase_Close_Fail(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The went into panic")
		}
	}()
	db, _ := NewDatabase("mysql", "")
	db.Close()
}