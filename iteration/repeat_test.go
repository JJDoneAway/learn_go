package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	length := len(Repeat("a", 12))
	expectedLength := 12

	if length != expectedLength {
		t.Errorf("Expected %d but was %d", expectedLength, length)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 12)
	}
}

func BenchmarkRepeat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat2("a", 12)
	}
}

func ExampleRepeat() {
	result := Repeat("b", 3)
	fmt.Println(result)
	//Output: bbb
}

func ExampleRepeat2() {
	result := Repeat2("b", 3)
	fmt.Println(result)
	//Output: bbb
}
