/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package sources

import (
	"encoding/json"
	"fmt"

	"github.com/devicechain-io/dc-event-sources/model"
	"github.com/rs/zerolog/log"
)

const (
	DECODER_TYPE_JSON = "json"
)

// Interface implemented by all decoders.
type Decoder interface {
	// Decodes a binary payload into an event.
	Decode(payload []byte) (*model.Event, error)
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

// Decode a json payload into an event.
func (jd *JsonDecoder) Decode(payload []byte) (*model.Event, error) {
	event := &model.Event{}
	err := json.Unmarshal(payload, event)
	if err != nil {
		return nil, err
	}
	log.Info().Msg(fmt.Sprintf("Parsed JSON event:\n\n%+v\n", event))
	return event, nil
}
