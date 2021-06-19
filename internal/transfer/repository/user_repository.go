package repository

import (
	"errors"

	"github.com/joubertredrat/walletlang/internal/transfer/entity"
)

var (
	UserRepositoryPayerNotFoundError = errors.New("Payer not found in user repository")
	UserRepositoryPayeeNotFoundError = errors.New("Payee not found in user repository")
)

type UserRepository interface {
	GetPayerByID(ID string) (*entity.Payer, error)
	GetPayeeByID(ID string) (*entity.Payee, error)
	UpdatePayer(payer *entity.Payer) error
	UpdatePayee(payee *entity.Payee) error
}
