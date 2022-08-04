package sum

import "testing"

func TestSum(t *testing.T) {
	array := [3]int{2, 2, 1}
	sum := Sum(array)
	expected := 5

	if sum != expected {
		t.Errorf("expected %q, but got %q", expected, sum)
	}
}
