package pointers

import (
	"errors"
	"fmt"
)

type String interface {
	String() string
}

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() (balance Bitcoin) {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func Divide(nominator, denominator float64) (float64, error) {

	if denominator == 0 {
		return 0, errors.New("division by zero")
	}

	return nominator / denominator, nil

}
