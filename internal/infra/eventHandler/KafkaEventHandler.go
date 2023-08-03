package eventHandler

import (
	"github.com/IBM/sarama"
	"github.com/ferminhg/learning-go/internal/domain"
	"log"
)

type KafkaEventHandler struct {
	adCollector sarama.SyncProducer
}

func NewKafkaEventHandler(brokerList []string) *KafkaEventHandler {
	return &KafkaEventHandler{
		adCollector: newAdCollector(brokerList),
	}
}

func newAdCollector(brokerList []string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	return producer
}

func (k KafkaEventHandler) Close() error {
	if err := k.adCollector.Close(); err != nil {
		log.Println("Failed to shut down data collector cleanly", err)
	}

	return nil
}

func (k KafkaEventHandler) SendMessage(msg *domain.ProducerMessage) (partition int32, offset int64, err error) {
	return k.adCollector.SendMessage(&sarama.ProducerMessage{
		Topic: msg.Topic(),
		Value: sarama.StringEncoder(msg.Value()),
	})
}
