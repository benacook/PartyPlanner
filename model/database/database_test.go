package database

import "testing"

//======================================================================================
func TestNewDatabase_Fail(t *testing.T) {
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
	NewMock()
	Db.Close()
}

//======================================================================================
//TestDatabase_Close_Fail causes panic.
func TestDatabase_Close_Fail(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The went into panic")
		}
	}()
	db, _ := NewDatabase("", "")
	db.Close()
}