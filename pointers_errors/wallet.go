package wallets

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(quantity Bitcoin) {
	w.balance += quantity
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var InsufficientBalanceError = errors.New("Unable to withdraw: insufficient balance")

func (w *Wallet) Withdraw(quantity Bitcoin) error {
	if quantity > w.balance {
		return InsufficientBalanceError
	}

	w.balance -= quantity
	return nil
}
