package event

import "github.com/joubertredrat/walletlang/internal/transfer/entity"

type TransactionEvent interface {
	DispatchWasScheduled(transaction *entity.Transaction) error
}
