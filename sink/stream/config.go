package stream

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// kafka producer-specific configuration
type Config struct {
	ProducerConfigMap *kafka.ConfigMap // kafka config map to be used by the producer
	Heartbeat int // stream interval in seconds
}