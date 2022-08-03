package iteration

import "testing"

func TestIterate(t *testing.T) {
	output := Iterate("a")
	expected := "aaa"

	if output != expected {
		t.Errorf("Expected %q, but got %q", expected, output)
	}
}
