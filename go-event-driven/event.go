// @author Shahrukh

// Purpose:
// The event.go file defines the Event struct representing a domain event.
// It includes a NewEvent function to create a new event with the specified aggregate ID, type, and data.

package main

import "time"

type Event struct {
	AggregateID string
	Type        string
	Data        interface{}
	Timestamp   time.Time
}

func NewEvent(aggregateID, eventType string, data interface{}) Event {
	return Event{
		AggregateID: aggregateID,
		Type:        eventType,
		Data:        data,
		Timestamp:   time.Now(),
	}
}
