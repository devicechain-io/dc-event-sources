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
	NewAssignment EventType = iota
	Location
	Measurement
	Alert
	StateChange
	CommandInvocation
	CommandResponse
)

var EventTypesByName map[string]EventType

// Unresolved event details.
type UnresolvedEvent struct {
	Source        string
	AltId         *string
	Device        string
	Assignment    *string
	OccurredTime  time.Time
	ProcessedTime time.Time
	EventType     EventType
	Payload       interface{}
}

// Payload for creating a new device assignment.
type NewAssignmentPayload struct {
	DeactivateExisting bool
	DeviceGroup        *string
	Asset              *string
	AssetGroup         *string
	Customer           *string
	CustomerGroup      *string
	Area               *string
	AreaGroup          *string
}

// Payload creating a new location.
type LocationPayload struct {
	Latitude  *string
	Longitude *string
	Elevation *string
}

// Initializer.
func init() {
	EventTypesByName = make(map[string]EventType)
	EventTypesByName[NewAssignment.String()] = NewAssignment
	EventTypesByName[Location.String()] = Location
	EventTypesByName[Measurement.String()] = Measurement
	EventTypesByName[Alert.String()] = Alert
	EventTypesByName[StateChange.String()] = StateChange
	EventTypesByName[CommandInvocation.String()] = CommandInvocation
	EventTypesByName[CommandResponse.String()] = CommandResponse
}
