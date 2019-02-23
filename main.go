package main

import "fmt"

func main() {
	account1 := &BankAccount{"1", 1000, "Jacob"}
	fmt.Println(account1)
	printOverdrafted(account1)
	account1.withdraw(1500)
	fmt.Println(account1)
	printOverdrafted(account1)
	account1.deposit(2000)
	fmt.Println(account1)
	printOverdrafted(account1)
}

func printOverdrafted(account *BankAccount) {
	if account.isOverdrafted() {
		fmt.Println("The account is overdrafted.")
	} else {
		fmt.Println("The account is not overdrafted.")
	}
}
