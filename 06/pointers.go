package pointers

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

// in Go, when you call a function or a method, the arguments are copied.
/*func (w *Wallet) Deposit(amount int) {
	w.balance += amount
} so hier, the w is a copy of whatever we called the method from
*/
func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Printf("address of balance in Deposit() is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufFunds = errors.New("cannot withdrawm insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufFunds
	}
	w.balance -= amount
	return nil
}

// The Stringer interface is defined in the fmt package and lets you define
// how your type is printed when used with the %s format string in prints.
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
