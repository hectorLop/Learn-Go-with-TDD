package Iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

// testing.B gives access to the cryptically named b.N
// the benchmark runs b.N times and measures how long it takes
func BenchmarkRepeat(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}