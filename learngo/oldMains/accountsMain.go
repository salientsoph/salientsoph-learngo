//package main
package oldMains

import (
	"fmt"

	"github.com/salientsoph/learngo/accounts"
)

func accountsMain() {
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

	/*
		account.Deposit(10) //여기서의 account == 복사본
		fmt.Println(account.Balance())
	*/

	/*
		go에서는 error를 직접 처리해줘야 한다(exception 등 없음..)
	*/

	/*
		err := account.Withdraw(20)
		if err != nil {
			log.Fatalln(err) //log.Fatalln: Println()을 호출하고 프로그램을 종료시킴
		}
	*/
	fmt.Println(account.Balance(), account.Owner())
}
