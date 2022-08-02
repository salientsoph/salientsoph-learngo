package accounts

// Account struct (make a comment)
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account //새로운 object를 return (actual memory address)
}

/*method

  func (reciever) methodName(parameter)

  reciever: struct의 첫글자를 따서 소문자로 짓는다
*/
// deposit x amount on your account
func (a Account) Deposit(amount int) {
	a.balance += amount
}

// balance of your account
func (a Account) Balance() int {
	return a.balance
}
