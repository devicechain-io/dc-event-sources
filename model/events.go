/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	"time"
)

type EventType int64

// Enumeration of event types.
//go:generate stringer -type=EventType
const (
	Location EventType = iota
	Measurement
	Alert
	StateChange
	CommandInvocation
	CommandResponse
)

var EventTypesByName map[string]EventType

// Base type for events.
type Event struct {
	Source        string
	AltId         *string
	Device        string
	Assignment    *string
	Customer      *string
	Area          *string
	Asset         *string
	OccurredTime  time.Time
	ProcessedTime time.Time
	EventType     EventType
}

// Payload for location event.
type LocationPayload struct {
	Latitude  *string
	Longitude *string
	Elevation *string
}

// Initializer.
func init() {
	EventTypesByName = make(map[string]EventType)
	EventTypesByName[Location.String()] = Location
	EventTypesByName[Measurement.String()] = Measurement
	EventTypesByName[Alert.String()] = Alert
	EventTypesByName[StateChange.String()] = StateChange
	EventTypesByName[CommandInvocation.String()] = CommandInvocation
	EventTypesByName[CommandResponse.String()] = CommandResponse
}
