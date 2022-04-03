/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package sources

import (
	"context"
	"strconv"

	"github.com/devicechain-io/dc-microservice/core"
	"github.com/rs/zerolog/log"
)

const (
	TYPE_MQTT = "mqtt"
)

type MqttEventSource struct {
	BrokerHost string
	BrokerPort int

	lifecycle core.LifecycleManager
}

// Create a new MQTT event source based on the given configuration.
func NewMqttEventSource(config map[string]string) (*MqttEventSource, error) {
	port, err := strconv.Atoi(config["port"])
	if err != nil {
		return nil, err
	}

	es := &MqttEventSource{
		BrokerHost: config["host"],
		BrokerPort: port,
	}
	es.lifecycle = core.NewLifecycleManager("mqtt-event-source", es, core.NewNoOpLifecycleCallbacks())
	return es, nil
}

// Initialize event source
func (es *MqttEventSource) Initialize(ctx context.Context) error {
	return es.lifecycle.Initialize(ctx)
}

// Initialize event source (as called by lifecycle manager)
func (es *MqttEventSource) ExecuteInitialize(ctx context.Context) error {
	log.Info().Msg("MQTT event source initialized.")
	return nil
}

// Start event source
func (es *MqttEventSource) Start(ctx context.Context) error {
	return es.lifecycle.Start(ctx)
}

// Start event source (as called by lifecycle manager)
func (es *MqttEventSource) ExecuteStart(ctx context.Context) error {
	log.Info().Msg("MQTT event source started.")
	return nil
}

// Stop event source
func (es *MqttEventSource) Stop(ctx context.Context) error {
	return es.lifecycle.Stop(ctx)
}

// Stop event source (as called by lifecycle manager)
func (es *MqttEventSource) ExecuteStop(ctx context.Context) error {
	log.Info().Msg("MQTT event source stopped.")
	return nil
}

// Terminate microservice
func (es *MqttEventSource) Terminate(ctx context.Context) error {
	return es.lifecycle.Terminate(ctx)
}

// Terminate event source (as called by lifecycle manager)
func (es *MqttEventSource) ExecuteTerminate(ctx context.Context) error {
	log.Info().Msg("MQTT event source terminated.")
	return nil
}
