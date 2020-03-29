package main

import "testing"

func TestPidor(t *testing.T) {
	expected := pidor()
	actual := "Ty pidor"
	if expected != actual {
		t.Fatalf("Wrong!")
	}
}
