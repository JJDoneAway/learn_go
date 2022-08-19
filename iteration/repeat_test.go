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

func ExampleSimpleLoop() {
	i := SimpleLoop()
	fmt.Println(i)
	// Output: 10
}

func ExampleNormalLoop() {
	i := NormalLoop()
	fmt.Println(i)
	// Output: 9

}

func ExampleCollectionLoop() {
	i := CollectionLoop([]int{4, 5, 6, 7})
	fmt.Println(i)
	//Output: 7
}

func ExampleMapLoop() {
	k, v := MapLoop(map[string]int{"Hallo": 1, "Pimmel": 2})
	fmt.Println(k, v)
	//Output: Pimmel 2
}

func ExampleIfCondition() {
	res := IfCondition()
	fmt.Println(res)
	//Output: true
}
