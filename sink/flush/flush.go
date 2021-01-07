package flush

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gitlab.com/idoko/shikari/db"
	"gitlab.com/idoko/shikari/models"
	"log"
)

var dbInstance db.Database

func Flush(ctx context.Context, cfg Config) error {
	dbInstance = cfg.DB

	cnsm, err := kafka.NewConsumer(cfg.ConsumerConfigMap)
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
				err := processMessage(e)
				if err != nil {
					log.Println("error while converting message to tweet: ", err.Error())
					continue
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

func processMessage(msg *kafka.Message) error {
	var tweet models.Tweet
	err := json.Unmarshal(msg.Value, &tweet)
	if err != nil {
		return err
	}

	err = dbInstance.SaveTweet(&tweet)
	if err != nil {
		return err
	}
	log.Println(fmt.Sprintf("Saved tweet: %s", tweet.TweetId))
	if msg.Headers != nil {
		fmt.Printf("%% Headers: %v\n", msg.Headers)
	}
	return nil
}

