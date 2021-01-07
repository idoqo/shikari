package stream

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gitlab.com/idoko/shikari/models"
	"io/ioutil"
	"log"
	"time"
)

func Stream(ctx context.Context, config Config) error {
	hits, err := loadJson()
	if err != nil {
		return err
	}

	prod, err := kafka.NewProducer(config.ProducerConfigMap)
	if err != nil {
		return err
	}

	doneChan := make(chan bool)
	go func() {
		defer close(doneChan)
		for e := range prod.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}
				return

			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}()

	topic := "shikari-stream"

	for _, tweet := range hits.Data {
		data, err := json.Marshal(tweet)
		if err != nil {
			// maybe log this, 'continue' the loop and move on.
			return err
		}
		// the select loop makes it possible to cancel execution without
		// waiting for all the 'tweets' currently in memory to be streamed i.e
		// when a ctrl+c is received (or ctx is cancelled), it streams the current tweet to kafka and halt.
		select {
		case <-ctx.Done():
			prod.Close()
			log.Println("quitting streamer hohoho")
			return nil
		default:
			prod.Produce(
				&kafka.Message{
					TopicPartition: kafka.TopicPartition{
						Topic: &topic, Partition: kafka.PartitionAny,
					},
			Value: data,
			}, nil)
			time.Sleep(1 * time.Second)
		}
	}
	_ = <-doneChan
	prod.Close()
	return nil
}

func loadJson() (models.SearchHits, error) {
	content, err := ioutil.ReadFile("junk/results1.json")
	hits := models.SearchHits{}

	if err != nil {
		return hits, err
	}

	err = json.Unmarshal(content, &hits)
	if err != nil {
		return hits, err
	}
	return hits, nil
}