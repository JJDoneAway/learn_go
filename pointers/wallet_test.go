package pointers

import (
	"fmt"
	"math"
	"testing"
)

func assertBalance(got, want Bitcoin, t *testing.T) {
	if want != got {
		t.Errorf("Got %d, but want %d.", got, want)
	}
}

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	fmt.Printf("Address of my balance %v\n", &wallet.balance)

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	want := Bitcoin(10)

	assertBalance(got, want, t)
	if got.String() != "10 BTC" {
		t.Errorf("Wrong format %s", got.String())
	}
}

func TestWithdraw(t *testing.T) {

	t.Run("withdraw positive", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		assertBalance(got, want, t)
	})

	t.Run("withdraw positive", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(30))
		got := wallet.Balance()
		want := Bitcoin(20)
		assertBalance(got, want, t)

		if err == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if err != ErrInsufficientFunds {
			t.Errorf("wanted %q but got %q", ErrInsufficientFunds, err)
		}
	})
}

func TestDivide(t *testing.T) {

	type testElement struct {
		nominator, denominator, expect float64
		fails                          bool
	}

	assert := func(test testElement) {
		got, err := Divide(test.nominator, test.denominator)

		if test.fails && err == nil {
			t.Fatal("Must fail but didn't", test)
		}

		if !test.fails && err != nil {
			t.Fatal("Must not fail but did", test, err)
		}

		if got != test.expect {
			t.Errorf("Expect %f but got %f", test.expect, got)
		}

	}

	assert(testElement{10.0, 100.0, 0.1, false})
	assert(testElement{10.0, 0, 0, true})
	assert(testElement{1, math.SmallestNonzeroFloat64, math.Inf(1), false})

}
