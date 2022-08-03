package main

import "testing"

func TestHello(t *testing.T) {
	actual := Hello("Rich")
	expected := "Hello Rich!"

	if actual != expected {
		t.Errorf("expected %q, actual %q", expected, actual)
	}
}
