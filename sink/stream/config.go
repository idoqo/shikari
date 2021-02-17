package stream

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

// kafka producer-specific configuration
type Config struct {
	ProducerConfigMap *kafka.ConfigMap // kafka config map to be used by the producer
	Heartbeat time.Duration // stream interval in seconds
}