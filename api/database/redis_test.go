package database

import "testing"

func TestCreateClient(t *testing.T) {
	r1 := CreateClient(1)
	r3 := CreateClient(3)
	
	if r1 == nil{
		t.Errorf("Redis not connected")
	}
	
	if r1 == nil{
		t.Errorf("Redis not connected")
	}
	
	if r1 == r3{
		t.Errorf("Different instances must be unique")
	}
}