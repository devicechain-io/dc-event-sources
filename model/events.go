/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	"math/big"
	"time"

	"github.com/google/uuid"
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
	Id            uuid.UUID
	AltId         *string
	Device        string
	Assignment    *uuid.UUID
	Customer      *string
	Area          *string
	Asset         *string
	OccuredTime   time.Time
	ProcessedTime time.Time
	EventType     EventType
}

// Event that contains location information.
type LocationEvent struct {
	Event
	Payload LocationPayload
}

// Payload for location event.
type LocationPayload struct {
	Latitude  big.Float
	Longitude big.Float
	Elevation big.Float
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
