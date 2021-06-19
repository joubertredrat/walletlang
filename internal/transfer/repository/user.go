package repository

import (
	"errors"

	"github.com/joubertredrat/walletlang/internal/transfer/entity"
)

var (
	PayerNotFoundError = errors.New("Payer not found in user repository")
	PayeeNotFoundError = errors.New("Payee not found in user repository")
)

type UserRepository interface {
	GetPayerByID(ID string) (*entity.Payer, error)
	GetPayeeByID(ID string) (*entity.Payee, error)
}
