package pointers_and_errors

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}
type Bitcoin int

type Stringer interface {
	String() string
}

var ErrInsufficientFunds = errors.New("oh no, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit method
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

// Balance method
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(bitcoin Bitcoin) error {
	if bitcoin > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= bitcoin
	return nil
}
