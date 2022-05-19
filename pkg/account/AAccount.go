package account

type AAccount struct {
	DefaultAccount
}

func NewAAccount() *AAccount {
	return &AAccount{*NewDefaultAccount()}
}
