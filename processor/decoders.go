/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package processor

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/devicechain-io/dc-event-sources/model"
)

const (
	DECODER_TYPE_JSON = "json"
)

// Payload expected for events passed in json format.
type JsonEvent struct {
	AltId        *string                `json:"altId,omitempty"`
	Device       string                 `json:"device"`
	Assignment   *string                `json:"assignment,omitempty"`
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

// Builds a new assignment payload from the json content.
func (jd *JsonDecoder) BuildNewAssignmentPayload(source *JsonEvent) (*model.NewAssignmentPayload, error) {
	payload := &model.NewAssignmentPayload{}

	// Any value works as true, but assume false if not passed.
	if _, ok := source.Payload["deactivateExisting"]; ok {
		payload.DeactivateExisting = true
	} else {
		payload.DeactivateExisting = false
	}

	if dgroup, ok := source.Payload["deviceGroup"]; ok {
		str := fmt.Sprintf("%v", dgroup)
		payload.DeviceGroup = &str
	}
	if asset, ok := source.Payload["asset"]; ok {
		str := fmt.Sprintf("%v", asset)
		payload.Asset = &str
	}
	if agroup, ok := source.Payload["assetGroup"]; ok {
		str := fmt.Sprintf("%v", agroup)
		payload.AssetGroup = &str
	}
	if cust, ok := source.Payload["customer"]; ok {
		str := fmt.Sprintf("%v", cust)
		payload.Customer = &str
	}
	if cgroup, ok := source.Payload["customerGroup"]; ok {
		str := fmt.Sprintf("%v", cgroup)
		payload.CustomerGroup = &str
	}
	if area, ok := source.Payload["area"]; ok {
		str := fmt.Sprintf("%v", area)
		payload.Customer = &str
	}
	if agroup, ok := source.Payload["areaGroup"]; ok {
		str := fmt.Sprintf("%v", agroup)
		payload.CustomerGroup = &str
	}

	return payload, nil
}

// Parses a locations event.
func (jd *JsonDecoder) BuildLocationsPayload(source *JsonEvent) (*model.LocationsPayload, error) {
	locbytes, err := json.Marshal(source.Payload)
	if err != nil {
		return nil, err
	}
	payload := &model.LocationsPayload{}
	json.Unmarshal(locbytes, payload)
	return payload, nil
}

// Parses a measurements event.
func (jd *JsonDecoder) BuildMeasurementsPayload(source *JsonEvent) (*model.MeasurementsPayload, error) {
	locbytes, err := json.Marshal(source.Payload)
	if err != nil {
		return nil, err
	}
	payload := &model.MeasurementsPayload{}
	json.Unmarshal(locbytes, payload)
	return payload, nil
}

// Parses an alerts event.
func (jd *JsonDecoder) BuildAlertsPayload(source *JsonEvent) (*model.AlertsPayload, error) {
	locbytes, err := json.Marshal(source.Payload)
	if err != nil {
		return nil, err
	}
	payload := &model.AlertsPayload{}
	json.Unmarshal(locbytes, payload)
	return payload, nil
}

// Parse json event payload.
func (jd *JsonDecoder) ParseEvent(payload []byte) (*JsonEvent, error) {
	jevent := &JsonEvent{}
	err := json.Unmarshal(payload, jevent)
	if err != nil {
		return nil, err
	}
	return jevent, nil
}

// Assemble an event based on json event data.
func (jd *JsonDecoder) AssembleEvent(jevent *JsonEvent) (*model.UnresolvedEvent, error) {
	event := &model.UnresolvedEvent{
		AltId:      jevent.AltId,
		Device:     jevent.Device,
		Assignment: jevent.Assignment,
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
	case model.NewAssignment:
		payload, err := jd.BuildNewAssignmentPayload(jevent)
		if err != nil {
			return nil, nil, err
		}
		return event, payload, nil
	case model.Location:
		payload, err := jd.BuildLocationsPayload(jevent)
		if err != nil {
			return nil, nil, err
		}
		return event, payload, nil
	case model.Measurement:
		payload, err := jd.BuildMeasurementsPayload(jevent)
		if err != nil {
			return nil, nil, err
		}
		return event, payload, nil
	case model.Alert:
		payload, err := jd.BuildAlertsPayload(jevent)
		if err != nil {
			return nil, nil, err
		}
		return event, payload, nil
	}

	return nil, nil, fmt.Errorf("unhandled event type: %s", jevent.EventType)
}
