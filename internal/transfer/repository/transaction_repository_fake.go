package repository

import "github.com/joubertredrat/walletlang/internal/transfer/entity"

type TransactionRepositoryFake struct {
	FakeCreate  func(transaction entity.Transaction) (*entity.Transaction, error)
	FakeUpdate  func(transaction entity.Transaction) error
	FakeGetByID func(ID string) (*entity.Transaction, error)
}

func NewTransactionRepositoryFake() TransactionRepositoryFake {
	return TransactionRepositoryFake{}
}

func (r TransactionRepositoryFake) Create(transaction entity.Transaction) (*entity.Transaction, error) {
	return r.FakeCreate(transaction)
}

func (r TransactionRepositoryFake) Update(transaction entity.Transaction) error {
	return r.FakeUpdate(transaction)
}

func (r TransactionRepositoryFake) GetByID(ID string) (*entity.Transaction, error) {
	return r.FakeGetByID(ID)
}
