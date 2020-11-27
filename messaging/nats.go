package messaging

import (
	"bytes"
	"encoding/gob"

	"go-cqrs/model"
	"go-cqrs/util"
	"github.com/nats-io/go-nats"
)

// NatsEventStore struct
type NatsEventStore struct {
	nc               *nats.Conn
	woofSubscription *nats.Subscription
	woofChan         chan WoofMessage
}

// OpenConnection to connecto to NATS
func OpenConnection(url string) (*NatsEventStore, error) {
	nc, err := nats.Connect(url)
	util.FailOnError(err, "Failed to connect to NATS")

	return &NatsEventStore{nc: nc}, nil
}

func (eventStore *NatsEventStore) writeMessage(m Message) ([]byte, error) {
	buffer := bytes.Buffer{}
	err := gob.NewEncoder(&buffer).Encode(m)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (eventStore *NatsEventStore) readMessage(data []byte, m interface{}) error {
	buffer := bytes.Buffer{}
	buffer.Write(data)

	return gob.NewDecoder(&buffer).Decode(m)
}

// PublishWoofMessage to publish a woof
func (eventStore *NatsEventStore) PublishWoofMessage(woof model.Woof) error {
	woofMessage := WoofMessage{woof.ID, woof.Body, woof.CreatedAt}
	data, err := eventStore.writeMessage(&woofMessage)
	if err != nil {
		return err
	}
	return eventStore.nc.Publish(woofMessage.Key(), data)
}

func (eventStore *NatsEventStore) SubscribeWoofMessage() (<-chan WoofMessage, error) {
	woofMessage := WoofMessage{}
	eventStore.woofChan = make(chan WoofMessage, 64)
	woofChan := make(chan *nats.Msg, 64)
	var err error
	eventStore.woofSubscription, err = eventStore.nc.ChanSubscribe(woofMessage.Key(), woofChan)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			select {
			case msg := <-woofChan:
				eventStore.readMessage(msg.Data, &woofMessage)
				eventStore.woofChan <- woofMessage
			}
		}
	}()
	return eventStore.woofChan, nil
}

// Close implementation
func (eventStore *NatsEventStore) Close() {
	eventStore.nc.Close()

	if eventStore.woofSubscription != nil {
		eventStore.woofSubscription.Unsubscribe()
	}

	close(eventStore.woofChan)
}
