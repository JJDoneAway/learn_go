package structs

type Triangle struct {
	BaseLine float64
	Height   float64
	Area     float64
}

func (t Triangle) Calc() {
	t.Area = .5 * t.BaseLine * t.Height
}

func (t *Triangle) Calc2() {
	t.Area = .5 * t.BaseLine * t.Height
}

func Calc3(t Triangle) {
	t.Area = .5 * t.BaseLine * t.Height
}

func Calc4(t []float64) {
	t[2] = .5 * t[0] * t[1]
}

func Calc5(t [3]float64) {
	t[2] = .5 * t[0] * t[1]
}
