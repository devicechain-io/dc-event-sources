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

// Marshals a payload based on what is expected for the given event type.
func MarshalPayloadForEventType(etype model.EventType, payload interface{}) ([]byte, error) {
	switch etype {
	case model.Location:
		if locpayload, ok := payload.(model.LocationPayload); ok {
			// Convert location payload.
			pbloc := &PLocationPayload{
				Latitude:  &locpayload.Latitude,
				Longitude: &locpayload.Longitude,
				Elevation: &locpayload.Elevation,
			}
			// Marshal event to bytes.
			bytes, err := proto.Marshal(pbloc)
			if err != nil {
				return nil, err
			}
			return bytes, nil
		}
		return nil, nil
	default:
		return nil, fmt.Errorf("unable to marshal payload for event type: %s", etype.String())
	}
}

// Marshal an event to protobuf bytes.
func MarshalEventToProtobuf(source string, event *model.Event, payload interface{}) ([]byte, error) {
	plbytes, err := MarshalPayloadForEventType(event.EventType, payload)
	if err != nil {
		return nil, err
	}

	// Encode protobuf event.
	pbevent := &PInboundEvent{
		SourceId:      source,
		AltId:         event.AltId,
		Device:        event.Device,
		Assignment:    event.Assignment,
		Customer:      event.Customer,
		Area:          event.Area,
		Asset:         event.Asset,
		OccurredTime:  event.OccurredTime.Format(time.RFC3339),
		ProcessedTime: event.ProcessedTime.Format(time.RFC3339),
		EventType:     event.EventType.String(),
		Payload:       plbytes,
	}

	// Marshal event to bytes.
	bytes, err := proto.Marshal(pbevent)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
