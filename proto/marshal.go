/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package proto

import (
	"fmt"
	"time"

	"github.com/devicechain-io/dc-event-sources/model"
	"google.golang.org/protobuf/proto"
)

// Marshal payload for a location event.
func MarshalPayloadForLocationEvent(payload *model.LocationPayload) ([]byte, error) {
	pbloc := &PLocationPayload{
		Latitude:  payload.Latitude,
		Longitude: payload.Longitude,
		Elevation: payload.Elevation,
	}
	bytes, err := proto.Marshal(pbloc)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Unmarshal a payload into a location event.
func UnmarshalPayloadForLocationEvent(payload []byte) (*model.LocationPayload, error) {
	pbloc := &PLocationPayload{}
	err := proto.Unmarshal(payload, pbloc)
	if err != nil {
		return nil, err
	}
	return &model.LocationPayload{
		Latitude:  pbloc.Latitude,
		Longitude: pbloc.Longitude,
		Elevation: pbloc.Elevation,
	}, nil
}

// Marshals a payload based on what is expected for the given event type.
func MarshalPayloadForEventType(etype model.EventType, payload interface{}) ([]byte, error) {
	switch etype {
	case model.Location:
		if locpayload, ok := payload.(*model.LocationPayload); ok {
			return MarshalPayloadForLocationEvent(locpayload)
		}
		return nil, fmt.Errorf("invalid location payload: %+v", payload)
	default:
		return nil, fmt.Errorf("unable to marshal payload for event type: %s", etype.String())
	}
}

// Unmarshal payload based on event type.
func UnmarshalPayloadForEventType(etype model.EventType, payload []byte) (interface{}, error) {
	switch etype {
	case model.Location:
		return UnmarshalPayloadForLocationEvent(payload)
	default:
		return nil, fmt.Errorf("unable to unmarshal payload for event type: %s", etype.String())
	}
}

// Marshal an event to protobuf bytes.
func MarshalEvent(event *model.Event, payload interface{}) ([]byte, error) {
	plbytes, err := MarshalPayloadForEventType(event.EventType, payload)
	if err != nil {
		return nil, err
	}

	// Encode protobuf event.
	pbevent := &PInboundEvent{
		SourceId:      event.Source,
		AltId:         event.AltId,
		Device:        event.Device,
		Assignment:    event.Assignment,
		Customer:      event.Customer,
		Area:          event.Area,
		Asset:         event.Asset,
		OccurredTime:  event.OccurredTime.Format(time.RFC3339),
		ProcessedTime: event.ProcessedTime.Format(time.RFC3339),
		EventType:     int64(event.EventType),
		Payload:       plbytes,
	}

	// Marshal event to bytes.
	bytes, err := proto.Marshal(pbevent)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Unmarshal encoded event.
func UnmarshalEvent(encoded []byte) (*model.Event, interface{}, error) {
	pbevent := &PInboundEvent{}
	err := proto.Unmarshal(encoded, pbevent)
	if err != nil {
		return nil, nil, err
	}
	etype := model.EventType(pbevent.EventType)
	payload, err := UnmarshalPayloadForEventType(etype, pbevent.Payload)
	if err != nil {
		return nil, nil, err
	}

	occtime, err := time.Parse(time.RFC3339, pbevent.OccurredTime)
	if err != nil {
		return nil, nil, err
	}
	proctime, err := time.Parse(time.RFC3339, pbevent.ProcessedTime)
	if err != nil {
		return nil, nil, err
	}
	event := &model.Event{
		Source:        pbevent.SourceId,
		AltId:         pbevent.AltId,
		Device:        pbevent.Device,
		Assignment:    pbevent.Assignment,
		Customer:      pbevent.Customer,
		Area:          pbevent.Area,
		Asset:         pbevent.Asset,
		OccurredTime:  occtime,
		ProcessedTime: proctime,
	}
	return event, payload, nil
}
