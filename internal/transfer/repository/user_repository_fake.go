package repository

import "github.com/joubertredrat/walletlang/internal/transfer/entity"

type UserRepositoryFake struct {
	FakeGetPayerByID func(ID string) (*entity.Payer, error)
	FakeGetPayeeByID func(ID string) (*entity.Payee, error)
	FakeCreatePayer  func(payer *entity.Payer) (*entity.Payer, error)
	FakeCreatePayee  func(payee *entity.Payee) (*entity.Payee, error)
	FakeUpdatePayer  func(payer *entity.Payer) (*entity.Payer, error)
	FakeUpdatePayee  func(payee *entity.Payee) (*entity.Payee, error)
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

func (r UserRepositoryFake) CreatePayer(payer *entity.Payer) (*entity.Payer, error) {
	return r.FakeCreatePayer(payer)
}

func (r UserRepositoryFake) CreatePayee(payee *entity.Payee) (*entity.Payee, error) {
	return r.FakeCreatePayee(payee)
}

func (r UserRepositoryFake) UpdatePayer(payer *entity.Payer) (*entity.Payer, error) {
	return r.FakeUpdatePayer(payer)
}

func (r UserRepositoryFake) UpdatePayee(payee *entity.Payee) (*entity.Payee, error) {
	return r.FakeUpdatePayee(payee)
}
