package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	reptime := 7
	thingtorep := "b"
	repeated := Repeat(thingtorep, reptime)
	expected := "bbbbbbb"

	if repeated != expected {
		t.Errorf("repeated %q expected %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// The testing.B gives you access to the cryptically named b.N
	for i := 0; i < b.N; i++ {
		Repeat("b", 7)
	}
	// When the benchmark is executed, it runs b.N times
	//and measures how long it takes.
}

func ExampleRepeat() {
	var char, count = "c", 9
	rep := Repeat(char, count)
	fmt.Println(rep)
	// Output: ccccccccc
}
