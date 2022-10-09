package main

import "fmt"

type Account struct {
	ID          string
	AccountType string
}

func (account *Account) create(accountType string) *Account {
	fmt.Println("account creation with type")
	account.AccountType = accountType
	return account
}

func (account *Account) getByID(id string) *Account {
	fmt.Println("getting account by ID")
	return account
}
func (account *Account) deletetByID(id string) {
	fmt.Println("deleting account by ID")
}

type Customer struct {
	name string
	id   string
}

func (customer *Customer) create(name string) *Customer {
	fmt.Println("Creating new Custoemr")
	customer.name = name
	return customer
}

type Transaction struct {
	id                       string
	amount                   float32
	sourceAccountNumber      string
	destinationAccountNumber string
}

func (transaction *Transaction) create(soureAccount string, destinatoinAccount string, amount float32) *Transaction {
	transaction.amount = amount
	transaction.destinationAccountNumber = destinatoinAccount
	transaction.sourceAccountNumber = soureAccount
	return transaction
}

type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

func (branchManager *BranchManagerFacade) creatAccount(accountName string, accountType string) (*Customer, *Account) {
	var account = branchManager.account.create(accountType)
	var customer = branchManager.customer.create(accountName)
	return customer, account
}

func (brachManager *BranchManagerFacade) createTransaction(srcAccountID string, destinationAccountID string, amount float32) *Transaction {
	var transaction = brachManager.transaction.create(srcAccountID, destinationAccountID, amount)
	return transaction
}

func main() {
	facade := NewBranchManagerFacade()
	var customer *Customer
	var account *Account
	var transaction *Transaction
	customer, account = facade.creatAccount("John Bellion", "Account 101")
	transaction = facade.createTransaction("account-frm", "account-to", 400000)
	fmt.Println("Cusonter name : ", customer.name)
	fmt.Println("Customers Account Type : ", account.AccountType)
	fmt.Println("Transacton between customters: ", transaction.amount)

}
