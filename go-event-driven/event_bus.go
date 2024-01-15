// @author Shahrukh

// Purpose:
// The event_bus.go file defines the EventBus interface for publishing events.
// It provides an implementation (ConsoleEventBus) of the EventBus interface, which prints the details of the event to the console.

package main

import "fmt"

type EventBus interface {
	PublishEvent(event Event)
}

type ConsoleEventBus struct{}

func NewConsoleEventBus() *ConsoleEventBus {
	return &ConsoleEventBus{}
}

func (bus *ConsoleEventBus) PublishEvent(event Event) {
	fmt.Printf("Event published: %+v\n", event)
}
