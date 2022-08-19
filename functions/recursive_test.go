package functions

import "testing"

func TestCa(t *testing.T) {
	res := Call()
	exp := "Hallo Welt!"

	if res != exp {
		t.Errorf("Got %s but want %s", res, exp)
	}
}

func TestCallCounter(t *testing.T) {
	res := CallCounter(1, 10)
	exp := 10

	if res != exp {
		t.Errorf("Want %d but got %d", exp, res)
	}
}

func TestFibonacci(t *testing.T) {

	got := CallFibonacci(100)
	exp := 144

	if got != exp {
		t.Errorf("Got %d but got %d", exp, got)
	}
}
