package main

import (
	"encoding/json"
	"fmt"
)

type AccountDetails struct {
	id          string
	accountType string
}

type Account struct {
	details      *AccountDetails
	CustomerName string
}

func (account *Account) setDetails(id, accountType string) {
	account.details = &AccountDetails{accountType: accountType, id: id}
}

func (account *Account) getID() string {
	return account.details.id
}
func (account *Account) getAccountType() string {
	return account.details.accountType
}
func main() {
	var account *Account = &Account{CustomerName: "Peter O"}
	jsonAccount, _ := json.Marshal(account)
	fmt.Println("Marshalled account: ", string(jsonAccount))
	account.setDetails("president", "government")
	fmt.Println("Account ID: ", account.getID())
	fmt.Println("Account Type: ", account.getAccountType())
}
