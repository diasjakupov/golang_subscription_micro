package messagequeue

import "log"

type MessageQueue interface {
	Publish(eventType string, data interface{}) error
}

type DummyMessageQueue struct{}

func (q *DummyMessageQueue) Publish(eventType string, data interface{}) error {
	log.Printf("Published event [%s]: %+v\n", eventType, data)
	return nil
}
