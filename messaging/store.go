package messaging

import (
	"go-cqrs/model"
)

// Message interface
type Message interface {
	Key() string
}

// EventStore interface
type EventStore interface {
	PublishWoofMessage(woof model.Woof) error
	SubscribeWoofMessage() (<-chan WoofMessage, error)
	// OnWoofMessageCreated(f func(WoofMessage)) error
	Close()
}

var impl EventStore

// SetEventStore implementation
func SetEventStore(eventStore EventStore) {
	impl = eventStore
}

// PublishWoofMessage implementation
func PublishWoofMessage(woof model.Woof) error {
	return impl.PublishWoofMessage(woof)
}

// SubscribeWoofMessage implementation
func SubscribeWoofMessage() (<-chan WoofMessage, error) {
	return impl.SubscribeWoofMessage()
}

// OnWoofMessageCreated implementation
// func OnWoofMessageCreated(f func(WoofMessage)) error {
// 	return impl.OnWoofMessageCreated(f)
// }

// Close implementation
func Close() {
	impl.Close()
}
