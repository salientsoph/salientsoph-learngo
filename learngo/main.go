package main

import (
	"fmt"

	"github.com/salientsoph/learngo/accounts"
)

func main() {
	/*
			만약

			banking.go
			type Account struct {
				owner   string
				balance int
			}

			main.go
			account := banking.Account{owner: "nicolas", balance: 1000}

			-> unknown field owner (error!)
			lowercase -> private하기에 접근 불가
					  -> 변수명을 uppercase로 바꿔준다(public)


		    banking.go
			type Account struct {
				Owner   string
				Balance int
			}

			main.go
			account := banking.Account{Owner: "nicolas", Balance: 1000}
			fmt.Println(account)
	*/

	account := accounts.NewAccount("nico")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())
}
