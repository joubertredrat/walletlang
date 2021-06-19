package repository

import (
	"errors"

	"github.com/joubertredrat/walletlang/internal/transfer/entity"
)

var (
	TransactionNotFoundError = errors.New("Transaction not found in transaction repository")
)

type TransactionRepository interface {
	Create(transaction entity.Transaction) (*entity.Transaction, error)
	GetByID(ID string) (*entity.Transaction, error)
}
