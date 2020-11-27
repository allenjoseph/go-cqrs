package messaging

import "time"

// WoofMessage struct
type WoofMessage struct {
	ID        string
	Body      string
	CreatedAt time.Time
}

// Key to publish a WoofMessage
func (m *WoofMessage) Key() string {
	return "woof.created"
}
