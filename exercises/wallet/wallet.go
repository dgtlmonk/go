package wallet

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds insufficient fund error
// The var keyword allows us to define values global to the package.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin usage Bitcoin(999)
type Bitcoin int

// Stringer ...
type Stringer interface {
	String() string
}

// String Bitcoin print format
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet ...
type Wallet struct {
	balance Bitcoin
}

// Deposit (pointer) so it doesn't copy
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance = amount
}

// Balance ...
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw ...
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		// returns error flag + string msg
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func wallet() {

}
