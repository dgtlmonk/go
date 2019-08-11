package wallet

import "fmt"

// Wallet ...
type Wallet struct {
	balance int
}

// Deposit (pointer) so it doesn't copy
func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)

	w.balance = amount
}

// Balance ...
func (w *Wallet) Balance() int {
	return w.balance
}

func wallet() {

}
