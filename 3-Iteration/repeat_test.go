package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 7)
	expected := "aaaaaaa"

	if repeated != expected {
		t.Errorf("Expected %s but got %s", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeatedChar := Repeat("w", 6)
	fmt.Println(repeatedChar)
	// Output: wwwwww
}
