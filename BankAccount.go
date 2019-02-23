package BankAccount

type BankAccount struct {
	id string
	money int
	customer string
}

func (BankAccount) deposit(amount int) int {
	money += amount
	return money
}

func (BankAccount) withdraw(amount int) int {
	money -= amount
	return money
}

func (BankAccount) isOverdrafted() bool {
	return money < 0
}