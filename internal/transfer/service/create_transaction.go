package service

import (
	"errors"

	"github.com/joubertredrat/walletlang/internal/transfer/entity"
	"github.com/joubertredrat/walletlang/internal/transfer/event"
	"github.com/joubertredrat/walletlang/internal/transfer/repository"
)

var (
	CreateTransactionHoustonError        = errors.New("Anything wrong is not right on create transaction")
	CreateTransactionPayerNotFoundError  = errors.New("Payer not found on create transaction")
	CreateTransactionPayeeNotFoundError  = errors.New("Payee not found on create transaction")
	CreateTransactionRepositoryError     = errors.New("Error to create transaction in repository on create transaction")
	CreateTransactionEventScheduledError = errors.New("Error to dispatch transaction scheduled in event on create transaction")
)

type CreateTransaction struct {
	UserRepository        repository.UserRepository
	TransactionRepository repository.TransactionRepository
	TransactionEvent      event.TransactionEvent
}

func NewCreateTransaction(
	u repository.UserRepository,
	t repository.TransactionRepository,
	e event.TransactionEvent,
) CreateTransaction {
	return CreateTransaction{
		UserRepository:        u,
		TransactionRepository: t,
		TransactionEvent:      e,
	}
}

func (c *CreateTransaction) HandleCreate(payerID, payeeID string, amount uint) (*entity.Transaction, error) {
	payer, err := c.UserRepository.GetPayerByID(payerID)
	if err != nil {
		if errors.Is(err, repository.UserRepositoryPayerNotFoundError) {
			return nil, CreateTransactionPayerNotFoundError
		}

		return nil, CreateTransactionHoustonError
	}

	payee, err := c.UserRepository.GetPayeeByID(payeeID)
	if err != nil {
		if errors.Is(err, repository.UserRepositoryPayeeNotFoundError) {
			return nil, CreateTransactionPayeeNotFoundError
		}

		return nil, CreateTransactionHoustonError
	}

	transaction, err := c.TransactionRepository.Create(entity.NewTransactionScheduled(*payer, *payee, amount))
	if err != nil {
		return nil, CreateTransactionRepositoryError
	}

	errEvent := c.TransactionEvent.DispatchWasScheduled(transaction)
	if errEvent != nil {
		return nil, CreateTransactionEventScheduledError
	}

	return transaction, nil
}
