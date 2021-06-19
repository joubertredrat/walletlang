package repository

import (
	"errors"

	"github.com/joubertredrat/walletlang/internal/transfer/entity"
)

type TransactionRepository interface {
	Create(transaction entity.Transaction) error
	GetByID(ID string) (*entity.Transaction, error)
}

var (
	TransactionNotFoundError = errors.New("Transaction not found")
)
