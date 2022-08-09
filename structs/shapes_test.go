package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	testTable := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 20.0}, 60.0},
		{Circle{4.0}, 25.13274122871834},
	}

	for _, value := range testTable {
		t.Helper()
		got := value.shape.Perimeter()
		if math.Abs(got-value.want) > 1e-7 {
			t.Errorf("Got %.2f want %.2f", got, value.want)
		}
	}

}

func TestArea(t *testing.T) {
	testTable := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.3, 24.54}, 301.842},
		{Circle{3.0}, 28.2743338},
	}

	for _, value := range testTable {
		t.Helper()
		got := value.shape.Area()
		if math.Abs(got-value.want) > 1e-7 {
			t.Errorf("Got %.2f want %.2f", got, value.want)
		}
	}
}
