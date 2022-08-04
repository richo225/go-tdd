package perimeter

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 5.0}

	output := Perimeter(10.0, 5.0)
	expected := 30.0

	if output != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, output)
	}
}

func TestArea(t *testing.T) {
	output := Area(20.0, 10.0)
	expected := 200.0

	if output != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, output)
	}
}
