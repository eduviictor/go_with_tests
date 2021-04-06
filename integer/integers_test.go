package integers

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d', result '%d'", expected, sum)
	}
}

func ExampleSum() {
	sum := Sum(1, 5)
	fmt.Println(sum)
	// Output: 6
}