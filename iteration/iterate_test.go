package iteration

import "testing"

func TestIterate(t *testing.T) {
	output := Iterate("a", 4)
	expected := "aaaa"

	if output != expected {
		t.Errorf("Expected %q, but got %q", expected, output)
	}
}

func BenchmarkIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("a", 3)
	}
}
