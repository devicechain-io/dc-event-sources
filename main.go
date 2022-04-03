/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"

	gql "github.com/graph-gophers/graphql-go"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"

	"github.com/devicechain-io/dc-event-sources/config"
	"github.com/devicechain-io/dc-event-sources/graphql"
	sources "github.com/devicechain-io/dc-event-sources/sources"
	"github.com/devicechain-io/dc-microservice/core"
	gqlcore "github.com/devicechain-io/dc-microservice/graphql"
	kcore "github.com/devicechain-io/dc-microservice/kafka"
)

var (
	Microservice  *core.Microservice
	Configuration *config.EventSourcesConfiguration
	EventSources  []core.LifecycleComponent

	GraphQLManager *gqlcore.GraphQLManager
	KakfaManager   *kcore.KafkaManager

	InboundEventsWriter *kafka.Writer
)

func main() {
	callbacks := core.LifecycleCallbacks{
		Initializer: core.LifecycleCallback{
			Preprocess:  func(context.Context) error { return nil },
			Postprocess: afterMicroserviceInitialized,
		},
		Starter: core.LifecycleCallback{
			Preprocess:  func(context.Context) error { return nil },
			Postprocess: afterMicroserviceStarted,
		},
		Stopper: core.LifecycleCallback{
			Preprocess:  beforeMicroserviceStopped,
			Postprocess: func(context.Context) error { return nil },
		},
		Terminator: core.LifecycleCallback{
			Preprocess:  beforeMicroserviceTerminated,
			Postprocess: func(context.Context) error { return nil },
		},
	}
	Microservice = core.NewMicroservice(callbacks)
	Microservice.Run()
}

// Parses the configuration from raw bytes.
func parseConfiguration() error {
	config := &config.EventSourcesConfiguration{}
	err := json.Unmarshal(Microservice.MicroserviceConfigurationRaw, config)
	if err != nil {
		return err
	}
	Configuration = config
	return nil
}

// Use configuration to build event sources.
func buildEventSources() error {
	created := make([]core.LifecycleComponent, 0)
	for _, source := range Configuration.EventSources {
		switch source.Type {
		case sources.TYPE_MQTT:
			mqtt, err := sources.NewMqttEventSource(source.Configuration)
			if err != nil {
				return err
			}
			created = append(created, mqtt)
		default:
			log.Warn().Msg(fmt.Sprintf("Unkown event source type: %s", source.Type))
		}
	}
	EventSources = created
	return nil
}

// Create kafka components used by this microservice.
func createKafkaComponents(kmgr *kcore.KafkaManager) error {
	ievents, err := kmgr.NewWriter(kmgr.NewScopedTopic(config.KAFKA_TOPIC_INBOUND_EVENTS))
	if err != nil {
		return err
	}
	InboundEventsWriter = ievents
	return nil
}

// Called after microservice has been initialized.
func afterMicroserviceInitialized(ctx context.Context) error {
	// Parse configuration.
	err := parseConfiguration()
	if err != nil {
		return err
	}

	// Build event sources from configuration.
	err = buildEventSources()
	if err != nil {
		return err
	}

	// Create and initialize kafka manager.
	KakfaManager = kcore.NewKafkaManager(Microservice, core.NewNoOpLifecycleCallbacks(), createKafkaComponents)
	err = KakfaManager.Initialize(ctx)
	if err != nil {
		return err
	}

	// Map of providers that will be injected into graphql http context.
	providers := map[gqlcore.ContextKey]interface{}{}

	// Create and initialize graphql manager.
	schema := gqlcore.CommonTypes + graphql.SchemaContent
	parsed := gql.MustParseSchema(schema, &graphql.SchemaResolver{})
	GraphQLManager = gqlcore.NewGraphQLManager(Microservice, core.NewNoOpLifecycleCallbacks(), *parsed, providers)
	err = GraphQLManager.Initialize(ctx)
	if err != nil {
		return err
	}

	// Initialize each event source.
	for _, source := range EventSources {
		err = source.Initialize(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// Called after microservice has been started.
func afterMicroserviceStarted(ctx context.Context) error {
	// Start kafka manager.
	err := KakfaManager.Start(ctx)
	if err != nil {
		return err
	}

	// Start graphql manager.
	err = GraphQLManager.Start(ctx)
	if err != nil {
		return err
	}

	// Start each event source.
	for _, source := range EventSources {
		err = source.Start(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// Called before microservice has been stopped.
func beforeMicroserviceStopped(ctx context.Context) error {
	// Stop each event source.
	for _, source := range EventSources {
		err := source.Stop(ctx)
		if err != nil {
			return err
		}
	}

	// Stop graphql manager.
	err := GraphQLManager.Stop(ctx)
	if err != nil {
		return err
	}

	// Stop kafka manager.
	err = KakfaManager.Stop(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Called before microservice has been terminated.
func beforeMicroserviceTerminated(ctx context.Context) error {
	// Terminate each event source.
	for _, source := range EventSources {
		err := source.Terminate(ctx)
		if err != nil {
			return err
		}
	}

	// Terminate graphql manager.
	err := GraphQLManager.Terminate(ctx)
	if err != nil {
		return err
	}

	// Terminate kafka manager.
	err = KakfaManager.Terminate(ctx)
	if err != nil {
		return err
	}

	return nil
}
