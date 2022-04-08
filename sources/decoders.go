/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package sources

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/devicechain-io/dc-event-sources/model"
	"github.com/rs/zerolog/log"
)

const (
	DECODER_TYPE_JSON = "json"
)

// Payload expected for events passed in json format.
type JsonEvent struct {
	AltId        *string                `json:"altId,omitempty"`
	Device       string                 `json:"device"`
	Assignment   *string                `json:"assignment,omitempty"`
	Customer     *string                `json:"customer,omitempty"`
	Area         *string                `json:"area,omitempty"`
	Asset        *string                `json:"asset,omitempty"`
	OccurredTime *string                `json:"occurredTime,omitempty"`
	EventType    string                 `json:"eventType"`
	Payload      map[string]interface{} `json:"payload"`
}

// Interface implemented by all decoders.
type Decoder interface {
	// Decodes a binary payload into an event.
	Decode(payload []byte) (*model.UnresolvedEvent, interface{}, error)
}

// Create a new decoder based on the given type indicator.
func NewDecoderForType(decodetype string, config map[string]string) (Decoder, error) {
	switch decodetype {
	case DECODER_TYPE_JSON:
		return NewJsonDecoder(config), nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Unknown decoder type: %s", decodetype))
	}
}

// Decodes payloads that use json format.
type JsonDecoder struct {
	Configuration map[string]string
}

// Create a new json decoder.
func NewJsonDecoder(config map[string]string) *JsonDecoder {
	return &JsonDecoder{
		Configuration: config,
	}
}

// Parses a location event.
func (jd *JsonDecoder) NewLocationPayload(source *JsonEvent) (*model.LocationPayload, error) {
	payload := &model.LocationPayload{}

	if latitude, ok := source.Payload["latitude"]; ok {
		latstr := fmt.Sprintf("%v", latitude)
		_, err := strconv.ParseFloat(latstr, 64)
		if err != nil {
			return nil, err
		}
		payload.Latitude = &latstr
	}
	if longitude, ok := source.Payload["longitude"]; ok {
		lonstr := fmt.Sprintf("%v", longitude)
		_, err := strconv.ParseFloat(lonstr, 64)
		if err != nil {
			return nil, err
		}
		payload.Longitude = &lonstr
	}
	if elevation, ok := source.Payload["elevation"]; ok {
		elestr := fmt.Sprintf("%v", elevation)
		_, err := strconv.ParseFloat(elestr, 64)
		if err != nil {
			return nil, err
		}
		payload.Elevation = &elestr
	}

	return payload, nil
}

// Parse json event payload.
func (jd *JsonDecoder) ParseEvent(payload []byte) (*JsonEvent, error) {
	jevent := &JsonEvent{}
	err := json.Unmarshal(payload, jevent)
	if err != nil {
		return nil, err
	}
	if log.Debug().Enabled() {
		log.Debug().Msg(fmt.Sprintf("Parsed JSON event:\n\n%+v\n", jevent))
	}
	return jevent, nil
}

// Assemble an event based on json event data.
func (jd *JsonDecoder) AssembleEvent(jevent *JsonEvent) (*model.UnresolvedEvent, error) {
	event := &model.UnresolvedEvent{
		AltId:      jevent.AltId,
		Device:     jevent.Device,
		Assignment: jevent.Assignment,
		Customer:   jevent.Customer,
		Area:       jevent.Area,
		Asset:      jevent.Asset,
	}
	if etype, ok := model.EventTypesByName[jevent.EventType]; ok {
		event.EventType = etype
	} else {
		return nil, fmt.Errorf("unknown event type in json payload: %s", jevent.EventType)
	}
	if jevent.OccurredTime != nil {
		otime, err := time.Parse(time.RFC3339, *jevent.OccurredTime)
		if err != nil {
			return nil, err
		}
		event.OccurredTime = otime
	} else {
		event.OccurredTime = time.Now()
	}
	event.ProcessedTime = time.Now()
	return event, nil
}

// Decode a json payload into an event.
func (jd *JsonDecoder) Decode(payload []byte) (*model.UnresolvedEvent, interface{}, error) {
	// Parse json payload.
	jevent, err := jd.ParseEvent(payload)
	if err != nil {
		return nil, nil, err
	}
	// Assemble event from json data.
	event, err := jd.AssembleEvent(jevent)
	if err != nil {
		return nil, nil, err
	}

	// Create payload based on event type.
	switch event.EventType {
	case model.Location:
		payload, err := jd.NewLocationPayload(jevent)
		if err != nil {
			return nil, nil, err
		}
		return event, payload, nil
	}

	return nil, nil, fmt.Errorf("unhandled event type: %s", jevent.EventType)
}
