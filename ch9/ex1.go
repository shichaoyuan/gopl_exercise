package main

import (
	"fmt"
)

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan int)
var withdrawResults = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	withdraws <- amount
	return <-withdrawResults
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case amount := <-withdraws:
			if balance < amount {
				withdrawResults <- false
			} else {
				balance -= amount
				withdrawResults <- true
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

func main() {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		ok := Withdraw(300)
		fmt.Println("withdraw :", ok)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		ok := Withdraw(100)
		fmt.Println("withdraw :", ok)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 200; got != want {
		fmt.Println("Balance = %d, want %d", got, want)
	}
}
