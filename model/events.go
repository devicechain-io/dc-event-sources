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
	Id            uuid.UUID              `json:"id"`
	AltId         string                 `json:"altId"`
	DeviceId      uint                   `json:"deviceId"`
	AssignmentId  uint                   `json:"assignmentId"`
	AreaId        uint                   `json:"areaId"`
	AssetId       uint                   `json:"assetId"`
	OccuredTime   time.Time              `json:"occuredTime"`
	ProcessedTime time.Time              `json:"processedTime"`
	EventType     EventType              `json:"eventType"`
	Payload       map[string]interface{} `json:"payload"`
}

// Payload for location event.
type LocationPayload struct {
	Latitude  big.Float `json:"latitude"`
	Longitude big.Float `json:"longitude"`
	Elevation big.Float `json:"elevation"`
}
