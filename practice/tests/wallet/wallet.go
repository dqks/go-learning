package wallet

import (
	"errors"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrInvalidAmount     = errors.New("invalid amount")
)

type Wallet struct {
	balance int
}

func NewWallet(initialBalance int) (*Wallet, error) {
	if initialBalance < 0 {
		return nil, ErrInvalidAmount
	}

	return &Wallet{balance: initialBalance}, nil
}

func (w *Wallet) Deposit(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	w.balance = w.balance + amount

	return nil
}

func (w *Wallet) Withdraw(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	result := w.balance - amount

	if result < 0 {
		return ErrInsufficientFunds
	}

	w.balance = result

	return nil
}

func (w *Wallet) Balance() int {
	return w.balance
}
