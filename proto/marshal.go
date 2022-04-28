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

// Marshal payload for a new assignment event.
func MarshalPayloadForNewAssignmentEvent(payload *model.NewAssignmentPayload) ([]byte, error) {
	pbna := &PNewAssignmentPayload{
		DeactivateExisting: payload.DeactivateExisting,
		DeviceGroup:        payload.DeviceGroup,
		Asset:              payload.Asset,
		AssetGroup:         payload.AssetGroup,
		Customer:           payload.Customer,
		CustomerGroup:      payload.CustomerGroup,
		Area:               payload.Area,
		AreaGroup:          payload.AreaGroup,
	}
	bytes, err := proto.Marshal(pbna)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Marshal payload for a location event.
func MarshalPayloadForLocationEvent(payload *model.LocationsPayload) ([]byte, error) {
	pbpayload := &PLocationsPayload{}
	for _, entry := range payload.Entries {
		pbentry := &PLocationEntry{
			Latitude:     entry.Latitude,
			Longitude:    entry.Longitude,
			Elevation:    entry.Elevation,
			OccurredTime: entry.OccurredTime,
		}
		pbpayload.Entries = append(pbpayload.Entries, pbentry)
	}
	bytes, err := proto.Marshal(pbpayload)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Unmarshal a payload into a new assignment event.
func UnmarshalPayloadForNewAssignmentEvent(payload []byte) (*model.NewAssignmentPayload, error) {
	pbassn := &PNewAssignmentPayload{}
	err := proto.Unmarshal(payload, pbassn)
	if err != nil {
		return nil, err
	}
	return &model.NewAssignmentPayload{
		DeactivateExisting: pbassn.DeactivateExisting,
		DeviceGroup:        pbassn.DeviceGroup,
		Asset:              pbassn.Asset,
		AssetGroup:         pbassn.AssetGroup,
		Customer:           pbassn.Customer,
		CustomerGroup:      pbassn.CustomerGroup,
		Area:               pbassn.Area,
		AreaGroup:          pbassn.AreaGroup,
	}, nil
}

// Unmarshal a payload into a location event.
func UnmarshalPayloadForLocationEvent(encoded []byte) (*model.LocationsPayload, error) {
	pbpayload := &PLocationsPayload{}
	err := proto.Unmarshal(encoded, pbpayload)
	if err != nil {
		return nil, err
	}
	payload := &model.LocationsPayload{}
	entries := make([]model.LocationEntry, 0)
	for _, pbentry := range pbpayload.Entries {
		entry := model.LocationEntry{
			Latitude:     pbentry.Latitude,
			Longitude:    pbentry.Longitude,
			Elevation:    pbentry.Elevation,
			OccurredTime: pbentry.OccurredTime,
		}
		entries = append(entries, entry)
	}
	payload.Entries = entries
	return payload, nil
}

// Marshals a payload based on what is expected for the given event type.
func MarshalPayloadForEventType(etype model.EventType, payload interface{}) ([]byte, error) {
	switch etype {
	case model.NewAssignment:
		if napayload, ok := payload.(*model.NewAssignmentPayload); ok {
			return MarshalPayloadForNewAssignmentEvent(napayload)
		}
		return nil, fmt.Errorf("invalid location payload: %+v", payload)
	case model.Location:
		if locpayload, ok := payload.(*model.LocationsPayload); ok {
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
	case model.NewAssignment:
		return UnmarshalPayloadForNewAssignmentEvent(payload)
	case model.Location:
		return UnmarshalPayloadForLocationEvent(payload)
	default:
		return nil, fmt.Errorf("unable to unmarshal payload for event type: %s", etype.String())
	}
}

// Marshal an unresolved event to protobuf bytes.
func MarshalUnresolvedEvent(event *model.UnresolvedEvent) ([]byte, error) {
	plbytes, err := MarshalPayloadForEventType(event.EventType, event.Payload)
	if err != nil {
		return nil, err
	}

	// Encode protobuf event.
	pbevent := &PUnresolvedEvent{
		SourceId:      event.Source,
		AltId:         event.AltId,
		Device:        event.Device,
		Assignment:    event.Assignment,
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

// Unmarshal encoded unresolved event.
func UnmarshalUnresolvedEvent(encoded []byte) (*model.UnresolvedEvent, error) {
	// Unmarshal protobuf event.
	pbevent := &PUnresolvedEvent{}
	err := proto.Unmarshal(encoded, pbevent)
	if err != nil {
		return nil, err
	}

	// Decode event type.
	etype := model.EventType(pbevent.EventType)

	// Unmarshal payload.
	payload, err := UnmarshalPayloadForEventType(etype, pbevent.Payload)
	if err != nil {
		return nil, err
	}

	occtime, err := time.Parse(time.RFC3339, pbevent.OccurredTime)
	if err != nil {
		return nil, err
	}
	proctime, err := time.Parse(time.RFC3339, pbevent.ProcessedTime)
	if err != nil {
		return nil, err
	}
	event := &model.UnresolvedEvent{
		Source:        pbevent.SourceId,
		AltId:         pbevent.AltId,
		Device:        pbevent.Device,
		Assignment:    pbevent.Assignment,
		OccurredTime:  occtime,
		ProcessedTime: proctime,
		EventType:     etype,
		Payload:       payload,
	}

	return event, nil
}
