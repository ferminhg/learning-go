package domain

type ProducerMessage struct {
	topic string
	value string //TODO: to remove
}

func (p ProducerMessage) Topic() string {
	return p.topic
}

func (p ProducerMessage) Value() string {
	return p.value
}

func NewProducerMessage(topic string, value string) *ProducerMessage {
	return &ProducerMessage{topic: topic, value: value}
}

type EventHandler interface {
	SendMessage(msg *ProducerMessage) (partition int32, offset int64, err error)
}
