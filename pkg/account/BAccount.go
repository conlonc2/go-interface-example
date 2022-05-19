package account

type BAccount struct {
	DefaultAccount
}

func NewBAccount() *BAccount {
	return &BAccount{*NewDefaultAccount()}
}
