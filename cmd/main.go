package main

import (
	"context"
	"fmt"
	"gitlab.com/idoko/shikari/sink"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go sink.StartStreaming(ctx, wg)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	log.Println(fmt.Sprint(<-sig))
	log.Println("Stopping Shikari...")

	cancel()
	wg.Wait()
}