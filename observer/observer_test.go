package observer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObserver_BroadcastNoSubscriber(t *testing.T) {
	pub := NewPublisher()
	result := pub.broadCast("Hello")
	assert.Equal(t, "Message Hello sent to 0 subscribers", result)
}

func TestObserver_BroadcastSubscriber(t *testing.T) {
	sub := newSubscriber("123")
	pub := NewPublisher()
	pub.addSubscriber(sub)
	result := pub.broadCast("Hello")
	assert.Equal(t, "Message Hello sent to 1 subscribers", result)
}
