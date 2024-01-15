// @author Shahrukh

// Purpose:
// The main.go file serves as the entry point of the application.
// It initializes an event store (eventStore) and an event bus (eventBus).
// Creates an aggregate with a unique ID and associates it with the event store and event bus.
// Applies initial events, simulates a command, retrieves events from the store, and prints the final state of the aggregate.

package main

import "fmt"

func main() {
	eventStore := NewEventStore()
	eventBus := NewConsoleEventBus()

	// Create an aggregate and apply events
	aggregateID := "123"
	aggregate := NewAggregate(aggregateID, eventStore, eventBus)
	aggregate.ApplyEvent(Event{AggregateID: aggregateID, Type: "ItemCreated", Data: "item-1"})

	// Simulate a command that results in events
	command := Command{AggregateID: aggregateID, Type: "UpdateItem", Data: "item-2"}
	aggregate.HandleCommand(command)

	// Read events from the event store and apply them
	events := eventStore.GetEvents(aggregateID)
	aggregate.ApplyEvents(events)

	// Print the final state of the aggregate
	fmt.Printf("Aggregate state after handling events: %+v\n", aggregate)
}
