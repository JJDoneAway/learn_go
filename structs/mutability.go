package structs

type AnyTriangle struct {
	BaseLine float64
	Height   float64
	Area     float64
}

func (t AnyTriangle) Calc() {
	t.Area = .5 * t.BaseLine * t.Height
}

func (t *AnyTriangle) Calc2() {
	t.Area = .5 * t.BaseLine * t.Height
}

func Calc3(t AnyTriangle) {
	t.Area = .5 * t.BaseLine * t.Height
}

func Calc4(t []float64) {
	t[2] = .5 * t[0] * t[1]
}

func Calc5(t [3]float64) {
	t[2] = .5 * t[0] * t[1]
}
