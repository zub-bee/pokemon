package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("c", 5)
	expected := "ccccc"

	if repeated != expected {
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 5))
	// Output: aaaaa
}
