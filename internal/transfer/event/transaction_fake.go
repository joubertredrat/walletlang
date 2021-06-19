package event

import "github.com/joubertredrat/walletlang/internal/transfer/entity"

type TransactionEventFake struct {
	FakeDispatchWasScheduled func(transaction *entity.Transaction) error
}

func NewTransactionEventFake() TransactionEventFake {
	return TransactionEventFake{}
}

func (e TransactionEventFake) DispatchWasScheduled(transaction *entity.Transaction) error {
	return e.FakeDispatchWasScheduled(transaction)
}
