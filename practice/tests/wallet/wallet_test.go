package wallet

import (
	"errors"
	"testing"
)

func TestNewWallet(t *testing.T) {
	tests := []struct {
		name    string
		balance int
		wantErr error
	}{
		{
			name:    "positive balance",
			balance: 1000,
			wantErr: nil,
		},
		{
			name:    "negative balance",
			balance: -1000,
			wantErr: ErrInvalidAmount,
		},
		{
			name:    "zero balance",
			balance: 0,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wallet, err := NewWallet(tt.balance)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf(
					"got %v but expected %v",
					err, tt.wantErr,
				)
			}

			if tt.wantErr == nil {
				if wallet.Balance() != tt.balance {
					t.Fatalf(
						"got balance %d but expected balance %d",
						wallet.Balance(),
						tt.balance,
					)
				}

			}

		})
	}
}

func TestDeposit(t *testing.T) {
	tests := []struct {
		name        string
		balance     int
		amount      int
		wantBalance int
		wantErr     error
	}{
		{
			name:        "positive amount",
			balance:     1000,
			amount:      300,
			wantBalance: 1300,
			wantErr:     nil,
		},
		{
			name:        "negative amount",
			balance:     1000,
			amount:      -300,
			wantBalance: 1000,
			wantErr:     ErrInvalidAmount,
		},
		{
			name:        "zero amount",
			balance:     1000,
			amount:      0,
			wantBalance: 1000,
			wantErr:     ErrInvalidAmount,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wallet, err := NewWallet(tt.balance)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got err %v when expected %v", err, tt.wantErr)
			}

			wallet.Deposit(tt.amount)

			if wallet.Balance() != tt.wantBalance {
				t.Fatalf(
					"got balance %d when deposit amount is %d instead of %d",
					wallet.Balance(),
					tt.amount,
					tt.wantBalance,
				)
			}

		})
	}

}

func TestWithdraw(t *testing.T) {
	tests := []struct {
		name        string
		balance     int
		amount      int
		wantBalance int
		wantErr     error
	}{
		{
			name:        "positive amount",
			balance:     1000,
			amount:      300,
			wantBalance: 700,
			wantErr:     nil,
		},
		{
			name:        "negative amount",
			balance:     1000,
			amount:      -300,
			wantBalance: 1000,
			wantErr:     ErrInvalidAmount,
		},
		{
			name:        "zero amount",
			balance:     1000,
			amount:      0,
			wantBalance: 1000,
			wantErr:     ErrInvalidAmount,
		},
		{
			name:        "full balance amount",
			balance:     1000,
			amount:      1000,
			wantBalance: 0,
			wantErr:     nil,
		},
		{
			name:        "more than balance amount",
			balance:     1000,
			amount:      1001,
			wantBalance: 1000,
			wantErr:     ErrInsufficientFunds,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wallet, err := NewWallet(tt.balance)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got err %v when expected %v", err, tt.wantErr)
			}

			wallet.Withdraw(tt.amount)

			if wallet.Balance() != tt.wantBalance {
				t.Fatalf(
					"got balance %d when deposit amount is %d instead of %d",
					wallet.Balance(),
					tt.amount,
					tt.wantBalance,
				)
			}
		})
	}
}
