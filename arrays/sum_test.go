package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		checkResult([]int{got}, []int{want}, t)

	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3, 4}, []int{3, 4}, []int{3}, []int{})
		want := []int{10, 7, 3, 0}

		checkResult(got, want, t)

	})
}

func BenchmarkSumAll(b *testing.B) {
	b.Run("Sum whole Slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SumAll([]int{1, 2, 3, 4}, []int{3, 4}, []int{3}, []int{})
		}
	})

	b.Run("Sum Tails", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SumAllTails([]int{1, 2, 3, 4}, []int{3, 4}, []int{3}, []int{})
		}
	})

}

func TestSumAllTails(t *testing.T) {
	t.Run("Sum arrays", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4}, []int{3, 4}, []int{3}, []int{})
		want := []int{9, 4, 0, 0}

		checkResult(got, want, t)

	})
}

func checkResult(got []int, want []int, t *testing.T) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want %v, but got %v", want, got)
	}
}

func TestRemoveElement(t *testing.T) {
	slc := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	fmt.Print(slc)

	mySlc := RemoveElement(2, slc)

	if mySlc[2] != "4" {
		t.Errorf("Expect 4 but got %v in %v", mySlc[2], mySlc)
	}
}
