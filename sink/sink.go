package sink

import (
	"context"
	"gitlab.com/idoko/shikari/sink/flush"
	"gitlab.com/idoko/shikari/sink/stream"
	"log"
	"sync"
)

type StreamConfig struct{

}

func StartStreaming(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			log.Println("stopping streamer")
			wg.Done()
			return
		default:
			// maybe this should be a case where a channel signal to the function
			// when it's time to check the data source again. That way, we only stream
			// to kafka at intervals.
			err := stream.Stream(ctx)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func StartFlushing(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			log.Println("stopping flusher...")
			wg.Done()
			return
		default:
			err := flush.Flush(ctx)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
