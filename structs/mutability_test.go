package structs

import "testing"

func TestMethod(t *testing.T) {

	t.Run("Immutable", func(t *testing.T) {
		r := AnyTriangle{4, 2, 0}
		r.Calc()

		if r.Area != 0 {
			t.Errorf("Got %g but want %g", r.Area, 4.0)
		}
	})

	t.Run("Mutable", func(t *testing.T) {
		r := AnyTriangle{4, 2, 0}
		r.Calc2()

		if r.Area != 4.0 {
			t.Errorf("Got %g but want %g", r.Area, 4.0)
		}
	})

}

func TestFunction(t *testing.T) {

	t.Run("Immutable", func(t *testing.T) {
		r := AnyTriangle{4, 2, 0}
		Calc3(r)

		if r.Area != 0 {
			t.Errorf("Got %g but want %g", r.Area, 4.0)
		}
	})

	t.Run("Mutable", func(t *testing.T) {
		r := []float64{4.0, 2.0, 0.0}
		Calc4(r)

		if r[2] != 4.0 {
			t.Errorf("Got %v but want %v", r, []float64{4.0, 2.0, 4.0})
		}
	})

	t.Run("Immutable Array", func(t *testing.T) {
		r := [3]float64{4.0, 2.0, 0.0}
		Calc5(r)

		if r[2] != 0.0 {
			t.Errorf("Got %v but want %v", r, [3]float64{4.0, 2.0, 0.0})
		}
	})
}
