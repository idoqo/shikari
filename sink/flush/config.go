package flush

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gitlab.com/idoko/shikari/db"
)

// Kafka consumer-specific configuration
type Config struct {
	ConsumerConfigMap *kafka.ConfigMap
	DB db.Database
}