package main

type BankAccount struct {
	id string
	gelt int
	customer string
}

func (b *BankAccount) deposit(amount int) int {
	b.gelt += amount
	return b.gelt
}

func (b *BankAccount) withdraw(amount int) int {
	b.gelt -= amount
	return b.gelt
}

func (b BankAccount) isOverdrafted() bool {
	return b.gelt < 0
}
