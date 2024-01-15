// @author Shahrukh

// Purpose:
// The event_store.go file defines the EventStore interface for storing and retrieving events.
// It provides an in-memory implementation (InMemoryEventStore) of the EventStore interface,
// allowing events to be saved and retrieved from an in-memory data structure.

package main

type EventStore interface {
	SaveEvent(event Event)
	GetEvents(aggregateID string) []Event
}

type InMemoryEventStore struct {
	Events map[string][]Event
}

func NewEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{
		Events: make(map[string][]Event),
	}
}

func (store *InMemoryEventStore) SaveEvent(event Event) {
	store.Events[event.AggregateID] = append(store.Events[event.AggregateID], event)
}

func (store *InMemoryEventStore) GetEvents(aggregateID string) []Event {
	return store.Events[aggregateID]
}
