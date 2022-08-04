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

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([3]int{1, 2, 3})
	}
	// Expected output 6
}
