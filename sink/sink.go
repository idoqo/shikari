package sink

import (
	"context"
	"gitlab.com/idoko/shikari/sink/flush"
	"gitlab.com/idoko/shikari/sink/stream"
	"log"
	"sync"
	"time"
)

func StartStreaming(ctx context.Context, wg *sync.WaitGroup, config stream.Config) {
	t := heartbeat(ctx, 30 * time.Second)

	for {
		select {
		case <-ctx.Done():
			log.Println("stopping streamer")
			wg.Done()
			return
		case <-t:
			err := stream.Stream(ctx, config)
			if err != nil {
				log.Println(err)
			}
		default:
		}
	}
}

// heartbeat is a time ticker with support for context cancellations.
func heartbeat(ctx context.Context, interval time.Duration) <- chan time.Time {
	notifier := make(chan time.Time, 1)
	// first run immediately (after 1 second), then run after `interval` subsequently
	startTime := time.Now().Add(1 * time.Second)

	// run in a goroutine, else the function blocks until the first event
	go func() {
		t := <-time.After(time.Until(startTime))
		notifier <- t

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case t2 := <-ticker.C:
				notifier <- t2
			case <-ctx.Done():
				close(notifier)
				return
			}
		}
	}()
	return notifier
}

func StartFlushing(ctx context.Context, wg *sync.WaitGroup, config flush.Config) {
	for {
		select {
		case <-ctx.Done():
			log.Println("stopping flusher...")
			wg.Done()
			return
		default:
			err := flush.Flush(ctx, config)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
