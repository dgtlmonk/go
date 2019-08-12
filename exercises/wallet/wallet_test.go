package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("wallet deposit test", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("wallet Withdraw test", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(5)}
		wallet.Withdraw(2)
		want := Bitcoin(3)

		assertBalance(t, wallet, want)
	})

	t.Run("wallet over-withdraw test", func(t *testing.T) {
		startingBalance := Bitcoin(20)

		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		// t.Fatal which will stop the test if it is called.
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
