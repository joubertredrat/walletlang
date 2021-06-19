package repository

import "github.com/joubertredrat/walletlang/internal/transfer/entity"

type UserRepositoryFake struct {
	FakeGetPayerByID func(ID string) (*entity.Payer, error)
	FakeGetPayeeByID func(ID string) (*entity.Payee, error)
	FakeUpdatePayer  func(payer *entity.Payer) error
	FakeUpdatePayee  func(payee *entity.Payee) error
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

func (r UserRepositoryFake) UpdatePayer(payer *entity.Payer) error {
	return r.FakeUpdatePayer(payer)
}

func (r UserRepositoryFake) UpdatePayee(payee *entity.Payee) error {
	return r.FakeUpdatePayee(payee)
}
