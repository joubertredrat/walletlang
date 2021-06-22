package event

import (
	"github.com/joubertredrat/walletlang/infrastructure/queue"
	eventEntity "github.com/joubertredrat/walletlang/infrastructure/transfer/event/entity"
	"github.com/joubertredrat/walletlang/internal/transfer/entity"
)

type TransactionEventKafka struct {
	producer queue.KafkaProducer
}

func NewTransactionEventKafka(producer queue.KafkaProducer) TransactionEventKafka {
	return TransactionEventKafka{
		producer: producer,
	}
}

func (e TransactionEventKafka) DispatchWasScheduled(transaction *entity.Transaction) error {
	transactionScheduled := eventEntity.NewTransactionScheduled(transaction.ID)

	return e.producer.Publish(transactionScheduled)
}
