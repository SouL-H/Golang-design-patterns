package main

import (
	"errors"
	"fmt"
)

type account struct {
	id string
}

func (a *account) checkAccount(id string) bool {
	fmt.Println("check account")
	return id == a.id
}

type security struct {
	code int
}

func (s *security) checkCode(code int) bool {
	fmt.Println("check code")
	return code == s.code
}

type wallet struct {
	amount int
}

func (w *wallet) credit(amount int) {
	fmt.Println("credit balance")
	w.amount += amount
}

func (w *wallet) debit(amount int) bool {
	if w.amount < amount {
		fmt.Println("debit balance failed")
		return false
	}
	fmt.Println("debit balance success")
	return true
}

type ledger struct{}

func (l *ledger) addEntry(accountID, txType string, amount int) {
	fmt.Printf("add ledger entry, account id: %s, tx type: %s, amount: %d\n",
		accountID, txType, amount)
}

type notification struct{}

func (n *notification) sendNotification() {
	fmt.Println("send notification")
}

type paymentFacade struct {
	account      *account
	security     *security
	wallet       *wallet
	ledger       *ledger
	notification *notification
}

func newPaymentFacade(accountID string, code int) *paymentFacade {
	fmt.Println("new payment facade")
	return &paymentFacade{
		account:      &account{accountID},
		security:     &security{code},
		wallet:       &wallet{},
		ledger:       &ledger{},
		notification: &notification{},
	}
}

func (p *paymentFacade) AddMoney(accountID string, securityCode, amount int) error {
	fmt.Println("add money")

	if ok := p.account.checkAccount(accountID); !ok {
		return errors.New("account is invalid")
	}

	if ok := p.security.checkCode(securityCode); !ok {
		return errors.New("code is invalid")
	}

	p.wallet.credit(amount)
	p.ledger.addEntry(accountID, "credit", amount)
	p.notification.sendNotification()
	return nil
}

func (p *paymentFacade) DeductMoney(accountID string, securityCode, amount int) error {
	fmt.Println("deduct money")

	if ok := p.account.checkAccount(accountID); !ok {
		return errors.New("account is invalid")
	}

	if ok := p.security.checkCode(securityCode); !ok {
		return errors.New("code is invalid")
	}

	p.wallet.debit(amount)
	p.ledger.addEntry(accountID, "debit", amount)
	p.notification.sendNotification()
	return nil
}

func main() {
	myAccount := "my_account"
	myCode := 9453

	f := newPaymentFacade(myAccount, myCode)
	if err := f.AddMoney(myAccount, myCode, 100); err != nil {
		fmt.Println("add money failed", err)
	}

	if err := f.DeductMoney(myAccount, myCode, 88); err != nil {
		fmt.Println("deduct money failed", err)
	}
}