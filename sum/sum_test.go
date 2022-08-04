package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("when passing collection of any size", func(t *testing.T) {
		array := []int{2, 2, 1, 4, 5}
		sum := Sum(array)
		expected := 14

		if sum != expected {
			t.Errorf("expected %q, but got %q", expected, sum)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3})
	}
	// Expected output 6
}

func TestSumAll(t *testing.T) {

	t.Run("when passing multiple slices", func(t *testing.T) {
		sum := SumAll([]int{2, 2}, []int{1, 1})
		expected := []int{4, 2}

		if !reflect.DeepEqual(sum, expected) {
			t.Errorf("expected %v, but got %v", expected, sum)
		}
	})
}
