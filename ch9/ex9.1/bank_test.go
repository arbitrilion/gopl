// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	bank "github.com/arbitrilion/gopl/ch9/ex9.1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("Alice deposit 200, balance =", bank.Balance())
		done <- struct{}{}
	}()

	// Alice
	go func() {
		bank.Withdraw(200)
		fmt.Println("Alice withdraw 200, balance =", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		fmt.Println("Bob deposit 100, balance =", bank.Balance())
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done
	<-done

	if got, want := bank.Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
