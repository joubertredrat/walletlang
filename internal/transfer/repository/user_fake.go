package repository

import "github.com/joubertredrat/walletlang/internal/transfer/entity"

type UserRepositoryFake struct {
	FakeGetPayerByID func(ID string) (*entity.Payer, error)
	FakeGetPayeeByID func(ID string) (*entity.Payee, error)
}

func NewUserRepositoryFake() UserRepositoryFake {
	return UserRepositoryFake{}
}

func (r UserRepositoryFake) GetPayerByID(ID string) (*entity.Payer, error) {
	return r.FakeGetPayerByID(ID)
}

func (r UserRepositoryFake) GetPayeeByID(ID string) (*entity.Payee, error) {
	return r.FakeGetPayeeByID(ID)
}
