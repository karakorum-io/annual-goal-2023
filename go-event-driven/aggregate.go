// @author Shahrukh

// Purpose:
// The aggregate.go file defines the Aggregate struct representing an entity that processes commands and applies events.
// It includes methods to apply events, handle commands, and update the state of the aggregate accordingly.

package main

type Command struct {
	AggregateID string
	Type        string
	Data        interface{}
}

type Aggregate struct {
	ID         string
	State      string
	Version    int
	EventStore EventStore
	EventBus   EventBus
}

func NewAggregate(id string, eventStore EventStore, eventBus EventBus) *Aggregate {
	return &Aggregate{
		ID:         id,
		EventStore: eventStore,
		EventBus:   eventBus,
	}
}

func (a *Aggregate) ApplyEvent(event Event) {
	// Apply business logic based on the event type
	switch event.Type {
	case "ItemCreated":
		a.State = event.Data.(string)
	case "ItemUpdated":
		a.State = event.Data.(string)
	}
	a.Version++
}

func (a *Aggregate) ApplyEvents(events []Event) {
	for _, event := range events {
		a.ApplyEvent(event)
	}
}

func (a *Aggregate) HandleCommand(command Command) {
	// Perform business logic based on the command type
	switch command.Type {
	case "UpdateItem":
		event := Event{AggregateID: command.AggregateID, Type: "ItemUpdated", Data: command.Data}
		a.EventStore.SaveEvent(event)
		a.EventBus.PublishEvent(event)
	}
}
