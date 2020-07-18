// ex9.1

package bank

import (
	"fmt"

	bank "github.com/arbitrilion/gopl/ch9/bank1"
)

var deposits = make(chan int)  // send amount to deposit
var withdraws = make(chan int) // send amount to withdraws
var balances = make(chan int)  // receive balance

// Deposit ...
func Deposit(amount int) { deposits <- amount }

// Withdraw ...
func Withdraw(amount int) { withdraws <- amount }

// Balance ...
func Balance() int { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
			fmt.Println("teller: deposit", amount, "succeeded, now balance = ", bank.Balance())
		case amount := <-withdraws:
			if amount <= balance {
				balance -= amount
				fmt.Println("teller: withdraw", amount, "succeeded, now balance = ", bank.Balance())
			} else {
				fmt.Println("teller: balance:", bank.Balance(), "not enough, withdraw", amount, "failed.")
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
