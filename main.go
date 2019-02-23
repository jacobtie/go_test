package main

import "fmt"

func main() {
	// Creates a pointer to a bank account object
	account1 := &BankAccount{"1", 1000, "Jacob"}
	// Prints the object
	fmt.Println(account1)
	// Prints whether or not the account is overdrawn by passing the reference
	printOverdrafted(account1)
	// Withdraws then prints the results
	fmt.Println("Withdrawing 1500")
	account1.withdraw(1500)
	fmt.Println(account1)
	printOverdrafted(account1)
	// Deposits then prints the results
	fmt.Println("Depositing 2000")
	account1.deposit(2000)
	fmt.Println(account1)
	printOverdrafted(account1)
	// Tests the return values of withdraw and deposit
	fmt.Println("Withdrawing 1000")
	fmt.Println(account1.withdraw(1000))
	fmt.Println("Depositing 500")
	fmt.Println(account1.deposit(500))
}

func printOverdrafted(account *BankAccount) {
	// Prints whether or not the account is overdrafted
	if account.isOverdrafted() {
		fmt.Println("The account is overdrafted.")
	} else {
		fmt.Println("The account is not overdrafted.")
	}
}
