package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"pikeys/internal/pikeys"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)

	service, err := pikeys.NewPiKeysService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	service.Run(group)

	eventChan := make(chan os.Signal, 1)
	signal.Notify(eventChan, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-eventChan:
	case <-ctx.Done():
	}
	cancel()
	if err := group.Wait(); err != nil {
		panic(err)
	}

}
