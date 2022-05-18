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
func MarshalPayloadForNewAssignmentEvent(payload *model.UnresolvedNewAssignmentPayload) ([]byte, error) {
	pbna := &PUnresolvedNewAssignmentPayload{
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

// Marshal payload for a locations event.
func MarshalPayloadForLocationsEvent(payload *model.UnresolvedLocationsPayload) ([]byte, error) {
	pbpayload := &PUnresolvedLocationsPayload{}
	for _, entry := range payload.Entries {
		pbentry := &PUnresolvedLocationEntry{
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

// Marshal payload for a measurements event.
func MarshalPayloadForMeasurementsEvent(payload *model.UnresolvedMeasurementsPayload) ([]byte, error) {
	pbpayload := &PUnresolvedMeasurementsPayload{}
	for _, entry := range payload.Entries {
		pbentry := &PUnresolvedMeasurementsEntry{
			Measurements: entry.Measurements,
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

// Marshal payload for an alerts event.
func MarshalPayloadForAlertsEvent(payload *model.UnresolvedAlertsPayload) ([]byte, error) {
	pbpayload := &PUnresolvedAlertsPayload{}
	for _, entry := range payload.Entries {
		pbentry := &PUnresolvedAlertEntry{
			Type:         entry.Type,
			Level:        entry.Level,
			Message:      entry.Message,
			Source:       entry.Source,
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
func UnmarshalPayloadForNewAssignmentEvent(payload []byte) (*model.UnresolvedNewAssignmentPayload, error) {
	pbassn := &PUnresolvedNewAssignmentPayload{}
	err := proto.Unmarshal(payload, pbassn)
	if err != nil {
		return nil, err
	}
	return &model.UnresolvedNewAssignmentPayload{
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

// Unmarshal a payload into a locations event.
func UnmarshalPayloadForLocationsEvent(encoded []byte) (*model.UnresolvedLocationsPayload, error) {
	pbpayload := &PUnresolvedLocationsPayload{}
	err := proto.Unmarshal(encoded, pbpayload)
	if err != nil {
		return nil, err
	}
	payload := &model.UnresolvedLocationsPayload{}
	entries := make([]model.UnresolvedLocationEntry, 0)
	for _, pbentry := range pbpayload.Entries {
		entry := model.UnresolvedLocationEntry{
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

// Unmarshal a payload into a measurements event.
func UnmarshalPayloadForMeasurementsEvent(encoded []byte) (*model.UnresolvedMeasurementsPayload, error) {
	pbpayload := &PUnresolvedMeasurementsPayload{}
	err := proto.Unmarshal(encoded, pbpayload)
	if err != nil {
		return nil, err
	}
	payload := &model.UnresolvedMeasurementsPayload{}
	entries := make([]model.UnresolvedMeasurementsEntry, 0)
	for _, pbentry := range pbpayload.Entries {
		entry := model.UnresolvedMeasurementsEntry{
			Measurements: pbentry.Measurements,
			OccurredTime: pbentry.OccurredTime,
		}
		entries = append(entries, entry)
	}
	payload.Entries = entries
	return payload, nil
}

// Unmarshal a payload into an alerts event.
func UnmarshalPayloadForAlertsEvent(encoded []byte) (*model.UnresolvedAlertsPayload, error) {
	pbpayload := &PUnresolvedAlertsPayload{}
	err := proto.Unmarshal(encoded, pbpayload)
	if err != nil {
		return nil, err
	}
	payload := &model.UnresolvedAlertsPayload{}
	entries := make([]model.UnresolvedAlertEntry, 0)
	for _, pbentry := range pbpayload.Entries {
		entry := model.UnresolvedAlertEntry{
			Type:         pbentry.Type,
			Level:        pbentry.Level,
			Message:      pbentry.Message,
			Source:       pbentry.Source,
			OccurredTime: pbentry.OccurredTime,
		}
		entries = append(entries, entry)
	}
	payload.Entries = entries
	return payload, nil
}

// Marshal unresolved payload based on event type.
func MarshalUnresolvedPayload(etype model.EventType, payload interface{}) ([]byte, error) {
	switch etype {
	case model.NewAssignment:
		if napayload, ok := payload.(*model.UnresolvedNewAssignmentPayload); ok {
			return MarshalPayloadForNewAssignmentEvent(napayload)
		}
		return nil, fmt.Errorf("invalid location payload: %+v", payload)
	case model.Location:
		if locpayload, ok := payload.(*model.UnresolvedLocationsPayload); ok {
			return MarshalPayloadForLocationsEvent(locpayload)
		}
		return nil, fmt.Errorf("invalid location payload: %+v", payload)
	case model.Measurement:
		if mxpayload, ok := payload.(*model.UnresolvedMeasurementsPayload); ok {
			return MarshalPayloadForMeasurementsEvent(mxpayload)
		}
		return nil, fmt.Errorf("invalid location payload: %+v", payload)
	case model.Alert:
		if apayload, ok := payload.(*model.UnresolvedAlertsPayload); ok {
			return MarshalPayloadForAlertsEvent(apayload)
		}
		return nil, fmt.Errorf("invalid location payload: %+v", payload)
	default:
		return nil, fmt.Errorf("unable to marshal unresolved payload for event type: %s", etype.String())
	}
}

// Unmarshal unresolved payload based on event type.
func UnmarshalUnresolvedPayload(etype model.EventType, payload []byte) (interface{}, error) {
	switch etype {
	case model.NewAssignment:
		return UnmarshalPayloadForNewAssignmentEvent(payload)
	case model.Location:
		return UnmarshalPayloadForLocationsEvent(payload)
	case model.Measurement:
		return UnmarshalPayloadForMeasurementsEvent(payload)
	case model.Alert:
		return UnmarshalPayloadForAlertsEvent(payload)
	default:
		return nil, fmt.Errorf("unable to unmarshal unresolved payload for event type: %s", etype.String())
	}
}

// Marshal an unresolved event to protobuf bytes.
func MarshalUnresolvedEvent(event *model.UnresolvedEvent) ([]byte, error) {
	plbytes, err := MarshalUnresolvedPayload(event.EventType, event.Payload)
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
	payload, err := UnmarshalUnresolvedPayload(etype, pbevent.Payload)
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
