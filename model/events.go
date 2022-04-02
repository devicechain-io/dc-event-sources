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

// Base type for events.
type Event struct {
	Id            uuid.UUID
	AltId         string
	DeviceId      uint
	AssignmentId  uint
	AreaId        uint
	AssetId       uint
	OccuredTime   time.Time
	ProcessedTime time.Time
	EventType     EventType
	Payload       interface{}
}

// Payload for location event.
type LocationPayload struct {
	Latitude  big.Float
	Longitude big.Float
	Elevation big.Float
}
