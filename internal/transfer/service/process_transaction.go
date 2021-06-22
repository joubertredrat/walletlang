package service

import (
	"errors"

	"github.com/joubertredrat/walletlang/internal/transfer/entity"
	"github.com/joubertredrat/walletlang/internal/transfer/repository"
)

var (
	ProcessTransactionHoustonError           = errors.New("Anything wrong is not right on process transaction")
	ProcessTransactionNotFoundError          = errors.New("Transaction not found on process transaction")
	ProcessTransactionPayerPayeeSameError    = errors.New("Payer and payee can not be same")
	ProcessTransactionInsufficientFundsError = errors.New("Insufficient payer funds on process transaction")
	ProcessTransactionWalletMovementError    = errors.New("Error on wallet movement on process transaction ")
)

type ProcessTransaction struct {
	UserRepository        repository.UserRepository
	TransactionRepository repository.TransactionRepository
}

func NewProcessTransaction(
	u repository.UserRepository,
	t repository.TransactionRepository,
) ProcessTransaction {
	return ProcessTransaction{
		UserRepository:        u,
		TransactionRepository: t,
	}
}

func (p *ProcessTransaction) HandleProcess(ID string) (*entity.Transaction, error) {
	transaction, err := p.TransactionRepository.GetByID(ID)
	if err != nil {
		return nil, ProcessTransactionNotFoundError
	}

	payer, err := p.UserRepository.GetPayerByID(transaction.Payer.ID)
	if err != nil {
		return nil, ProcessTransactionHoustonError
	}

	payee, err := p.UserRepository.GetPayeeByID(transaction.Payee.ID)
	if err != nil {
		return nil, ProcessTransactionHoustonError
	}

	if payer.ID == payee.ID {
		transaction.SetStatusPayerPayeeSame()
		p.TransactionRepository.Update(*transaction)
		return nil, ProcessTransactionPayerPayeeSameError
	}

	if transaction.Amount > payer.Amount {
		transaction.SetStatusInsufficientFunds()
		p.TransactionRepository.Update(*transaction)
		return nil, ProcessTransactionInsufficientFundsError
	}

	payer.Amount = payer.Amount - transaction.Amount
	payee.Amount = payee.Amount + transaction.Amount
	transaction.SetStatusDone()

	if _, err := p.UserRepository.UpdatePayer(payer); err != nil {
		return nil, ProcessTransactionWalletMovementError
	}
	if _, err := p.UserRepository.UpdatePayee(payee); err != nil {
		return nil, ProcessTransactionWalletMovementError
	}
	p.TransactionRepository.Update(*transaction)

	return transaction, nil
}
