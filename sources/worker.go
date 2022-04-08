/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package sources

import (
	"fmt"

	"github.com/devicechain-io/dc-event-sources/model"
	"github.com/rs/zerolog/log"
)

// Worker used to decode event payloads.
type DecodeWorker struct {
	WorkerId    int
	SourceId    string
	Decoder     Decoder
	RawMessages <-chan []byte
	Callback    func(string, *model.UnresolvedEvent, interface{})
	Failed      func(string, []byte, error)
}

// Create a new decode worker.
func NewDecodeWorker(workerId int, sourceId string, decoder Decoder, rawMessages <-chan []byte,
	callback func(string, *model.UnresolvedEvent, interface{}),
	failed func(string, []byte, error)) *DecodeWorker {
	worker := &DecodeWorker{
		WorkerId:    workerId,
		SourceId:    sourceId,
		Decoder:     decoder,
		RawMessages: rawMessages,
		Callback:    callback,
		Failed:      failed,
	}
	return worker
}

// Processes raw payloads into decoded events.
func (wrk *DecodeWorker) Process() {
	for {
		raw, more := <-wrk.RawMessages
		if more {
			log.Debug().Msg(fmt.Sprintf("Decode handled by worker id %d", wrk.WorkerId))
			event, payload, err := wrk.Decoder.Decode(raw)
			if err != nil {
				wrk.Failed(wrk.SourceId, raw, err)
			} else {
				wrk.Callback(wrk.SourceId, event, payload)
			}
		} else {
			log.Debug().Msg("Decode worker received shutdown signal.")
			return
		}
	}
}
