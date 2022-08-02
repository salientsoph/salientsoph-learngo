package accounts

import (
	"errors"
	"fmt"
)

// Account struct (make a comment)
type Account struct {
	owner   string
	balance int
}

//error 정의 시 err로 시작해야함
var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account //새로운 object를 return (actual memory address)
}

/*method

  func (reciever) methodName(parameter)

  reciever: struct의 첫글자를 따서 소문자로 짓는다

  *Account: pointer receiver
  -> go에서는 복사본을 receiver로 받기 때문에,
  복사본을 받지 않도록 pointer를 사용했다.

*/
// deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// balance of your account
func (a Account) Balance() int {
	return a.balance
}

// withdraw from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
	/** error

	boolean에 true, false 가 있는 것처럼
	error에는 error, nil(= None, null) 이 있음
	*/
}

// changeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

/* __str__

	python
	print(class) 하면 __str__을 호출함

	go
	func () String() string{

}


*/
func (a Account) String() string {
	//Sprint(): string 형태
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
