package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	expected := 3

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

func AddExample() {
	sum := Add(3, 4)
	fmt.Println(sum)
	// Output: 7
}
