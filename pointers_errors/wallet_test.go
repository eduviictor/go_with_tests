package wallets

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		confirmBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		confirmBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdrawing with sufficient balance", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		confirmBalance(t, wallet, Bitcoin(10))
		confirmNonExistentError(t, err)
	})

	t.Run("Withdrawing with insufficient balance", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{initialBalance}
		err := wallet.Withdraw(Bitcoin(100))

		confirmBalance(t, wallet, initialBalance)
		confirmError(t, err, InsufficientBalanceError)
	})
}

func confirmBalance(t *testing.T, wallet Wallet, expect Bitcoin) {
	t.Helper()
	result := wallet.Balance()

	if result != expect {
		t.Errorf("result %s, expect %s", result, expect)
	}
}

func confirmError(t *testing.T, result error, expect error) {
	t.Helper()

	if result == nil {
		t.Fatal("Expected an error but none occurred")
	}

	if result != expect {
		t.Errorf("result %s, expect %s", result, expect)
	}
}

func confirmNonExistentError(t *testing.T, result error) {
	t.Helper()

	if result != nil {
		t.Fatal("Unexpected error received")
	}
}
