/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package sources

import (
	"context"
	"fmt"
	"strconv"

	"github.com/devicechain-io/dc-event-sources/model"
	"github.com/devicechain-io/dc-microservice/core"
	"github.com/rs/zerolog/log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	TYPE_MQTT = "mqtt"
)

type MqttEventSource struct {
	BrokerHost string
	BrokerPort int
	Topic      string

	Client  mqtt.Client
	Decoder Decoder

	lifecycle core.LifecycleManager
	callback  func(*model.Event, interface{})
}

// Create a new MQTT event source based on the given configuration.
func NewMqttEventSource(config map[string]string, decoder Decoder, callback func(*model.Event, interface{})) (*MqttEventSource, error) {
	port, err := strconv.Atoi(config["port"])
	if err != nil {
		return nil, err
	}

	es := &MqttEventSource{
		BrokerHost: config["host"],
		BrokerPort: port,
		Topic:      config["topic"],
		Decoder:    decoder,
	}
	es.lifecycle = core.NewLifecycleManager("mqtt-event-source", es, core.NewNoOpLifecycleCallbacks())
	es.callback = callback
	return es, nil
}

// Called when message is received from topic.
func (es *MqttEventSource) onMessage(client mqtt.Client, msg mqtt.Message) {
	if log.Debug().Enabled() {
		log.Debug().Msg(fmt.Sprintf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic()))
	}
	event, payload, err := es.Decoder.Decode(msg.Payload())
	if err != nil {
		log.Error().Err(err).Msg("Unable to decode event message.")
		return
	}
	es.callback(event, payload)
}

// Called on successful connection.
func (es *MqttEventSource) onConnect(client mqtt.Client) {
	log.Info().Msg("MQTT event source connected successfully.")
}

// Called when connection is lost.
func (es *MqttEventSource) onConnectionLost(client mqtt.Client, err error) {
	log.Info().Msg("MQTT event source connection lost.")
}

// Initialize event source
func (es *MqttEventSource) Initialize(ctx context.Context) error {
	return es.lifecycle.Initialize(ctx)
}

// Initialize event source (as called by lifecycle manager)
func (es *MqttEventSource) ExecuteInitialize(ctx context.Context) error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", es.BrokerHost, es.BrokerPort))
	opts.SetClientID("devicechain")
	opts.SetDefaultPublishHandler(es.onMessage)
	opts.OnConnect = es.onConnect
	opts.OnConnectionLost = es.onConnectionLost
	es.Client = mqtt.NewClient(opts)
	if token := es.Client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	log.Info().Msg("MQTT event source initialized.")
	return nil
}

// Start event source
func (es *MqttEventSource) Start(ctx context.Context) error {
	return es.lifecycle.Start(ctx)
}

// Start event source (as called by lifecycle manager)
func (es *MqttEventSource) ExecuteStart(ctx context.Context) error {
	token := es.Client.Subscribe(es.Topic, 1, es.onMessage)
	token.Wait()
	log.Info().Msg(fmt.Sprintf("MQTT event source subscribed to topic '%s'.", es.Topic))
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
