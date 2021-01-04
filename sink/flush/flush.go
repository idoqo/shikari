package flush

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func Flush(ctx context.Context) error {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id": "shikari-consumers",
		"session.timeout.ms": 6000,
		"auto.offset.reset": "earliest",

	}
	cnsm, err := kafka.NewConsumer(configMap)
	if err != nil {
		return err
	}
	topic := "shikari-stream"
	err = cnsm.Subscribe(topic, nil)
	if err != nil {
		return err
	}

	run := true
	for run == true {
		select {
		case <- ctx.Done():
			log.Println("caught cancellation from context, terminating")
			run = false
		default:
			ev := cnsm.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
			case kafka.Error:
				// errors should be informational and the client will try to
				// recover, but halt if all brokers are down.
				log.Println("Error: ", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				log.Println("ignored...")
			}
		}
	}
	log.Println("closing consumer")
	cnsm.Close()
	return nil
}

