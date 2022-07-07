package observer

import "fmt"

type Publisher interface {
	addSubscriber(subscriber Subscriber)
	removeSubscriber(subId string)
	broadCast(msg string) string
}

type publisher struct {
	subscribers map[string]Subscriber
}

func NewPublisher() publisher {
	return publisher{
		subscribers: make(map[string]Subscriber),
	}
}

func (p publisher) addSubscriber(subscriber Subscriber) {
	p.subscribers[subscriber.getId()] = subscriber

}

func (p publisher) removeSubscriber(subId string) {
	delete(p.subscribers, subId)
}

func (p publisher) broadCast(msg string) string {
	count := 0
	for _, subscriber := range p.subscribers {
		subscriber.react(msg)
		count += 1
	}
	return fmt.Sprintf("Message %s sent to %d subscribers", msg, count)
}

type Subscriber interface {
	getId() string
	react(msg string) string
}

// subscriber implementation
type subscriber struct {
	subId string
}

func (s subscriber) getId() string {
	return s.subId
}

func (s subscriber) react(msg string) string {
	return fmt.Sprintf("ID: %s recieved message %s", s.subId, msg)
}

func newSubscriber(subId string) subscriber {
	return subscriber{
		subId: subId,
	}
}
