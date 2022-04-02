/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package config

const (
	KAFKA_TOPIC_INBOUND_EVENTS = "inbound-events"
)

type NestedConfiguration struct {
	Test string
}

type EventSourcesConfiguration struct {
	Nested NestedConfiguration
}

// Creates the default event sources configuration
func NewEventSourcesConfiguration() *EventSourcesConfiguration {
	return &EventSourcesConfiguration{
		Nested: NestedConfiguration{
			Test: "test",
		},
	}
}
