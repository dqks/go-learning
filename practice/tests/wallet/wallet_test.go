package wallet_test

import (
	"errors"
	"test-practice/wallet"
	"testing"
)

func TestNewWallet(t *testing.T) {
	tests := []struct {
		name       string
		balance    int
		wantErr    error
		wantWallet bool
	}{
		{
			name:       "positive balance",
			balance:    1000,
			wantErr:    nil,
			wantWallet: true,
		},
		{
			name:       "negative balance",
			balance:    -1000,
			wantErr:    wallet.ErrInvalidAmount,
			wantWallet: false,
		},
		{
			name:       "zero balance",
			balance:    0,
			wantErr:    nil,
			wantWallet: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := wallet.NewWallet(tt.balance)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf(
					"got %v but expected %v",
					err, tt.wantErr,
				)
			}

			if w != nil && !tt.wantWallet {
				t.Fatalf(
					"w is not nil when balance is %d",
					tt.balance,
				)
			} else if w == nil && tt.wantWallet {
				t.Fatalf(
					"w is nil when balance is %d",
					tt.balance,
				)
			}

			if tt.wantErr == nil {
				if w.Balance() != tt.balance {
					t.Fatalf(
						"got balance %d but expected balance %d",
						w.Balance(),
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
			wantErr:     wallet.ErrInvalidAmount,
		},
		{
			name:        "zero amount",
			balance:     1000,
			amount:      0,
			wantBalance: 1000,
			wantErr:     wallet.ErrInvalidAmount,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := wallet.NewWallet(tt.balance)

			if err != nil {
				t.Errorf("failed to create wallet %v", err)
			}

			err = w.Deposit(tt.amount)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got err %v when expected %v", err, tt.wantErr)
			}

			if w.Balance() != tt.wantBalance {
				t.Fatalf(
					"got balance %d when deposit amount is %d instead of %d",
					w.Balance(),
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
			wantErr:     wallet.ErrInvalidAmount,
		},
		{
			name:        "zero amount",
			balance:     1000,
			amount:      0,
			wantBalance: 1000,
			wantErr:     wallet.ErrInvalidAmount,
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
			wantErr:     wallet.ErrInsufficientFunds,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := wallet.NewWallet(tt.balance)

			if err != nil {
				t.Errorf("failed to create wallet %v", err)
			}

			err = w.Withdraw(tt.amount)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got err %v when expected %v", err, tt.wantErr)
			}

			if w.Balance() != tt.wantBalance {
				t.Fatalf(
					"got balance %d when withdraw amount is %d instead of %d",
					w.Balance(),
					tt.amount,
					tt.wantBalance,
				)
			}
		})
	}
}

func TestTransfer(t *testing.T) {
	tests := []struct {
		name            string
		balanceFrom     int
		balanceTo       int
		amount          int
		wantErr         error
		wantBalanceFrom int
		wantBalanceTo   int
	}{
		{
			name:            "positive amount",
			balanceFrom:     1000,
			balanceTo:       0,
			amount:          200,
			wantErr:         nil,
			wantBalanceFrom: 800,
			wantBalanceTo:   200,
		},
		{
			name:            "negative amount",
			balanceFrom:     1000,
			balanceTo:       0,
			amount:          -1,
			wantErr:         wallet.ErrInvalidAmount,
			wantBalanceFrom: 1000,
			wantBalanceTo:   0,
		},
		{
			name:            "zero amount",
			balanceFrom:     1000,
			balanceTo:       0,
			amount:          0,
			wantErr:         wallet.ErrInvalidAmount,
			wantBalanceFrom: 1000,
			wantBalanceTo:   0,
		},
		{
			name:            "amount more than balanceFrom",
			balanceFrom:     1000,
			balanceTo:       0,
			amount:          1001,
			wantErr:         wallet.ErrInsufficientFunds,
			wantBalanceFrom: 1000,
			wantBalanceTo:   0,
		},
		{
			name:            "amount less than balanceFrom",
			balanceFrom:     1000,
			balanceTo:       0,
			amount:          999,
			wantErr:         nil,
			wantBalanceFrom: 1,
			wantBalanceTo:   999,
		},
		{
			name:            "amount equals balanceFrom",
			balanceFrom:     1000,
			balanceTo:       0,
			amount:          1000,
			wantErr:         nil,
			wantBalanceFrom: 0,
			wantBalanceTo:   1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			walletFrom, err := wallet.NewWallet(tt.balanceFrom)

			if err != nil {
				t.Fatalf(
					"failed to create walletFrom %v with balance %d",
					err,
					tt.balanceFrom,
				)
			}

			walletTo, err := wallet.NewWallet(tt.balanceTo)

			if err != nil {
				t.Fatalf(
					"failed to create walletTo %v with balance %d",
					err,
					tt.balanceTo,
				)
			}

			err = walletFrom.Transfer(walletTo, tt.amount)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf(
					"got  %v expected %v",
					err,
					tt.wantErr,
				)
			}

			if walletFrom.Balance() != tt.wantBalanceFrom {
				t.Fatalf(
					"got result balanceFrom %d expected %d",
					walletFrom.Balance(),
					tt.balanceFrom,
				)
			}

			if walletTo.Balance() != tt.wantBalanceTo {
				t.Fatalf(
					"got result balanceTo %d expected %d",
					walletTo.Balance(),
					tt.wantBalanceTo,
				)
			}

		})
	}
}
