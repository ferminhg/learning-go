package eventHandler

import (
	"github.com/IBM/sarama"
	"github.com/IBM/sarama/mocks"
	"github.com/ferminhg/learning-go/internal/domain"
)

type MockEventHandler struct {
	MockSP *mocks.SyncProducer
}

func NewMockEventHandler(t mocks.ErrorReporter) MockEventHandler {

	sp := mocks.NewSyncProducer(t, nil)
	defer func() {
		if err := sp.Close(); err != nil {
			t.Errorf(err.Error())
		}
	}()

	return MockEventHandler{sp}
}

func (m MockEventHandler) SendMessage(msg *domain.ProducerMessage) (partition int32, offset int64, err error) {
	pm := &sarama.ProducerMessage{
		Topic: msg.Topic(),
		Value: sarama.StringEncoder(msg.Value()),
	}

	return m.MockSP.SendMessage(pm)
}
