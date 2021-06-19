package service_test

import (
	"errors"
	"testing"

	"github.com/joubertredrat/walletlang/internal/transfer/entity"
	"github.com/joubertredrat/walletlang/internal/transfer/event"
	"github.com/joubertredrat/walletlang/internal/transfer/repository"
	"github.com/joubertredrat/walletlang/internal/transfer/service"
	"github.com/stretchr/testify/assert"
)

func getCreateTransactionUserRepositoryFake() repository.UserRepository {
	repo := repository.NewUserRepositoryFake()
	repo.FakeGetPayerByID = func(ID string) (*entity.Payer, error) {
		switch ID {
		case "60f70370-7316-41d5-8980-7ebb9ebb7e68", "0accc07e-ea9a-4b7d-bf99-2060aa2c46bf":
			return nil, repository.UserRepositoryPayerNotFoundError
		case "853ea721-09f7-41f9-a6d4-01592f79d40c", "a7733d09-e989-499b-b705-d3dfbd5827f2":
			return nil, errors.New("fail to connect into fake mongodb")
		default:
			payer := entity.NewPayer(ID)
			return &payer, nil
		}
	}
	repo.FakeGetPayeeByID = func(ID string) (*entity.Payee, error) {
		switch ID {
		case "cc5c676d-9363-4c39-b48b-472fd26ff961", "a698ccf1-a416-44ae-bf86-b2ecf23c5763":
			return nil, repository.UserRepositoryPayeeNotFoundError
		case "6e4f5873-2833-4ea2-b03e-7ce034fbad1a", "a6d8d6c1-0ec7-46ea-b1a6-e02aac89717d":
			return nil, errors.New("fail to connect into fake mongodb")
		default:
			payee := entity.NewPayee(ID)
			return &payee, nil
		}
	}

	return repo
}

func TestCreateTransactionHandleCreate(t *testing.T) {
	tests := []struct {
		name                     string
		payerID                  string
		payeeID                  string
		amount                   uint
		getUserRepository        func() repository.UserRepository
		getTransactionRepository func() repository.TransactionRepository
		getTransactionEvent      func() event.TransactionEvent
		getTransactionExpected   func(PayerID, PayeeID string, Amount uint) *entity.Transaction
		getErrorExpected         func() error
	}{
		{
			name:              "Create transaction with success",
			payerID:           "45015f90-6285-45eb-be1e-90eeccb448bc",
			payeeID:           "f2934e46-055b-4846-828b-90d94728f0de",
			amount:            1250,
			getUserRepository: getCreateTransactionUserRepositoryFake,
			getTransactionRepository: func() repository.TransactionRepository {
				repo := repository.NewTransactionRepositoryFake()
				repo.FakeCreate = func(transaction entity.Transaction) (*entity.Transaction, error) {
					transaction.ID = "7004f709-5779-4472-80af-4b5338758a6c"
					return &transaction, nil
				}

				return repo
			},
			getTransactionEvent: func() event.TransactionEvent {
				event := event.NewTransactionEventFake()
				event.FakeDispatchWasScheduled = func(transaction *entity.Transaction) error {
					return nil
				}

				return event
			},
			getTransactionExpected: func(PayerID, PayeeID string, Amount uint) *entity.Transaction {
				payer := entity.NewPayer(PayerID)
				payee := entity.NewPayee(PayeeID)
				transaction := entity.NewTransactionScheduled(payer, payee, Amount)
				transaction.ID = "7004f709-5779-4472-80af-4b5338758a6c"
				return &transaction
			},
			getErrorExpected: func() error {
				return nil
			},
		},
		{
			name:              "Create transaction with payer not found",
			payerID:           "60f70370-7316-41d5-8980-7ebb9ebb7e68",
			payeeID:           "f2934e46-055b-4846-828b-90d94728f0de",
			amount:            1250,
			getUserRepository: getCreateTransactionUserRepositoryFake,
			getTransactionRepository: func() repository.TransactionRepository {
				return repository.NewTransactionRepositoryFake()
			},
			getTransactionEvent: func() event.TransactionEvent {
				return event.NewTransactionEventFake()
			},
			getTransactionExpected: func(PayerID, PayeeID string, Amount uint) *entity.Transaction {
				return nil
			},
			getErrorExpected: func() error {
				return service.CreateTransactionPayerNotFoundError
			},
		},
		{
			name:              "Create transaction with payee not found",
			payerID:           "45015f90-6285-45eb-be1e-90eeccb448bc",
			payeeID:           "cc5c676d-9363-4c39-b48b-472fd26ff961",
			amount:            1250,
			getUserRepository: getCreateTransactionUserRepositoryFake,
			getTransactionRepository: func() repository.TransactionRepository {
				return repository.NewTransactionRepositoryFake()
			},
			getTransactionEvent: func() event.TransactionEvent {
				return event.NewTransactionEventFake()
			},
			getTransactionExpected: func(PayerID, PayeeID string, Amount uint) *entity.Transaction {
				return nil
			},
			getErrorExpected: func() error {
				return service.CreateTransactionPayeeNotFoundError
			},
		},
		{
			name:              "Create transaction with user repository error to get payer",
			payerID:           "853ea721-09f7-41f9-a6d4-01592f79d40c",
			payeeID:           "f2934e46-055b-4846-828b-90d94728f0de",
			amount:            1250,
			getUserRepository: getCreateTransactionUserRepositoryFake,
			getTransactionRepository: func() repository.TransactionRepository {
				return repository.NewTransactionRepositoryFake()
			},
			getTransactionEvent: func() event.TransactionEvent {
				return event.NewTransactionEventFake()
			},
			getTransactionExpected: func(PayerID, PayeeID string, Amount uint) *entity.Transaction {
				return nil
			},
			getErrorExpected: func() error {
				return service.CreateTransactionHoustonError
			},
		},
		{
			name:              "Create transaction with user repository error to get payee",
			payerID:           "45015f90-6285-45eb-be1e-90eeccb448bc",
			payeeID:           "6e4f5873-2833-4ea2-b03e-7ce034fbad1a",
			amount:            1250,
			getUserRepository: getCreateTransactionUserRepositoryFake,
			getTransactionRepository: func() repository.TransactionRepository {
				return repository.NewTransactionRepositoryFake()
			},
			getTransactionEvent: func() event.TransactionEvent {
				return event.NewTransactionEventFake()
			},
			getTransactionExpected: func(PayerID, PayeeID string, Amount uint) *entity.Transaction {
				return nil
			},
			getErrorExpected: func() error {
				return service.CreateTransactionHoustonError
			},
		},
		{
			name:              "Create transaction with error on create transaction into repository",
			payerID:           "45015f90-6285-45eb-be1e-90eeccb448bc",
			payeeID:           "f2934e46-055b-4846-828b-90d94728f0de",
			amount:            1250,
			getUserRepository: getCreateTransactionUserRepositoryFake,
			getTransactionRepository: func() repository.TransactionRepository {
				repo := repository.NewTransactionRepositoryFake()
				repo.FakeCreate = func(transaction entity.Transaction) (*entity.Transaction, error) {
					return nil, errors.New("fail to connect into fake mongodb")
				}

				return repo
			},
			getTransactionEvent: func() event.TransactionEvent {
				return event.NewTransactionEventFake()
			},
			getTransactionExpected: func(PayerID, PayeeID string, Amount uint) *entity.Transaction {
				return nil
			},
			getErrorExpected: func() error {
				return service.CreateTransactionRepositoryError
			},
		},
		{
			name:              "Create transaction with error on dispatch scheduled into event",
			payerID:           "45015f90-6285-45eb-be1e-90eeccb448bc",
			payeeID:           "f2934e46-055b-4846-828b-90d94728f0de",
			amount:            1250,
			getUserRepository: getCreateTransactionUserRepositoryFake,
			getTransactionRepository: func() repository.TransactionRepository {
				repo := repository.NewTransactionRepositoryFake()
				repo.FakeCreate = func(transaction entity.Transaction) (*entity.Transaction, error) {
					transaction.ID = "7004f709-5779-4472-80af-4b5338758a6c"
					return &transaction, nil
				}

				return repo
			},
			getTransactionEvent: func() event.TransactionEvent {
				event := event.NewTransactionEventFake()
				event.FakeDispatchWasScheduled = func(transaction *entity.Transaction) error {
					return errors.New("fail to sent data to kafka cluster")
				}

				return event
			},
			getTransactionExpected: func(PayerID, PayeeID string, Amount uint) *entity.Transaction {
				return nil
			},
			getErrorExpected: func() error {
				return service.CreateTransactionEventScheduledError
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			transactionExpected := test.getTransactionExpected(test.payerID, test.payeeID, test.amount)
			errorExpected := test.getErrorExpected()

			createTransaction := service.NewCreateTransaction(
				test.getUserRepository(),
				test.getTransactionRepository(),
				test.getTransactionEvent(),
			)

			transactionGot, errorGot := createTransaction.HandleCreate(test.payerID, test.payeeID, test.amount)

			assert.Equal(t, transactionExpected, transactionGot)
			assert.Equal(t, errorExpected, errorGot)
		})
	}
}
