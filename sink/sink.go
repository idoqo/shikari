package sink

import (
	"context"
	"gitlab.com/idoko/shikari/sink/stream"
	"log"
	"sync"
)

func StartStreaming(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			log.Println("quitting")
			wg.Done()
			return
		default:
			err := stream.Stream(ctx)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
