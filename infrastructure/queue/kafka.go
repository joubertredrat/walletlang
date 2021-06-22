package queue

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"github.com/joubertredrat/walletlang/infrastructure/transfer/event/entity"
)

type KafkaProducer struct {
	brokerHost string
	topic      string
}

func NewKafkaProducer(brokerHost, topic string) KafkaProducer {
	return KafkaProducer{
		brokerHost: brokerHost,
		topic:      topic,
	}
}

func (k *KafkaProducer) Publish(transactionScheduled entity.TransactionScheduled) error {
	document, _ := json.Marshal(transactionScheduled)

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{k.brokerHost}, config)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Panic(err)
		}
	}()
	msg := &sarama.ProducerMessage{
		Topic: *&k.topic,
		Value: sarama.StringEncoder(document),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
		return err
	}

	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", k.topic, partition, offset)
	return nil
}
