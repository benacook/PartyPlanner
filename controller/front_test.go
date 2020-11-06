package controller

import "testing"

//theres not much to test here, so we just make sure it does not panic

//======================================================================================
func TestGetRouter(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("The went into panic")
		}
	}()
	GetRouter()
}

//======================================================================================
func TestRegisterHandlers(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("The went into panic")
		}
	}()
	RegisterHandlers()
}
