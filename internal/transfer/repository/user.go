package repository

import (
	"errors"

	"github.com/joubertredrat/walletlang/internal/transfer/entity"
)

type UserRepository interface {
	GetPayerByID(ID string) (*entity.Payer, error)
	GetPayeeByID(ID string) (*entity.Payee, error)
}

var (
	PayerNotFoundError = errors.New("Payer not found")
	PayeeNotFoundError = errors.New("Payee not found")
)
