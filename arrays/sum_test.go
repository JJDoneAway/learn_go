package arrays

import (
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
