package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repetitions := Repeat("a", 7)
	expected := "aaaaaaa"

	if repetitions != expected {
		t.Errorf("expected '%s', result '%s'", expected, repetitions)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	result := Repeat("b", 2)
	fmt.Println(result)
	// Output: bb
}