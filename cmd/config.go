package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gitlab.com/idoko/shikari/db"
	"gitlab.com/idoko/shikari/sink/flush"
	"gitlab.com/idoko/shikari/sink/stream"
	"log"
)

type Config struct {
	FlusherConfig flush.Config
	StreamerConfig stream.Config
	Logger *log.Logger // todo: replace calls to log.* with the config logger.
}

type ConfigOption func(cfg *Config)

func NewConfig(opts ...ConfigOption) Config {
	cfg := Config{
		FlusherConfig: flush.Config{},
		StreamerConfig: stream.Config{},
		Logger: nil,
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

func ConfigureFlusher(bootstrapServers string, timeOut int, groupId string, db db.Database) ConfigOption {
	return func(cfg *Config) {
		configMap := &kafka.ConfigMap{
			"bootstrap.servers": bootstrapServers,
			"session.timeout.ms": timeOut,
			"group.id": groupId,
			"auto.offset.reset": "earliest",
		}
		cfg.FlusherConfig = flush.Config{
			ConsumerConfigMap: configMap,
			DB: db,
		}
	}
}

func ConfigureStreamer(bootstrapServers string, heartbeat int) ConfigOption {
	return func(cfg *Config) {
		configMap := &kafka.ConfigMap{
			"bootstrap.servers": bootstrapServers,
		}
		cfg.StreamerConfig = stream.Config{
			ProducerConfigMap: configMap,
			Heartbeat: heartbeat,
		}
	}
}