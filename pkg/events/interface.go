package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Dispatch(event EventInterface)
	Register(eventName string, handler EventHandlerInterface) error // AddListener
	Remove(eventName string, handler EventHandlerInterface) error   // RemoveListener
	Has(eventName string) bool                                      // HasListener
	Clear() error                                                   // ClearListeners
}
