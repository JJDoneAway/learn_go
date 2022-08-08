package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkResult := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if math.Abs(got-want) > 1e-7 {
			t.Errorf("Got %.2f want %.2f", got, want)
		}
	}

	t.Run("Perimeter of a rect", func(t *testing.T) {
		want := 60.0
		checkResult(t, Rectangle{10.0, 20.0}, want)

	})

	t.Run("Perimeter of a circle", func(t *testing.T) {
		want := 25.132741228718346

		checkResult(t, Circle{4.0}, want)
	})

}

func TestArea(t *testing.T) {
	checkResult := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if math.Abs(got-want) > 1e-7 {
			t.Errorf("Got %.2f want %.2f", got, want)
		}
	}

	t.Run("Area of an Rect", func(t *testing.T) {
		want := 301.842

		checkResult(t, Rectangle{12.3, 24.54}, want)
	})

	t.Run("Area of a Circle", func(t *testing.T) {
		want := 28.2743338

		checkResult(t, Circle{3.0}, want)
	})

}
