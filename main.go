package main

import (
	"context"
	"time"

	"log"

	"github.com/matthagenbuch/followtheleader/physical"
	"github.com/matthagenbuch/followtheleader/spanner"
)

const (
	bucketName    = "matth-follow-the-leader"
	leadershipKey = "shard-0"
)

func main() {
	logger := log.Default()
	// b, err := gcs.NewBackend(map[string]string{
	// 	"bucket":     bucketName,
	// 	"ha_enabled": "true",
	// }, logger)
	b, err := spanner.NewBackend(map[string]string{
		"database":   "projects/lightstep-dev/instances/development/databases/dev-matth",
		"ha_table":   "alertevaluator_elections_ha",
		"ha_enabled": "true",
	}, logger)
	if err != nil {
		panic(err)
	}
	haBackend, ok := b.(physical.HABackend)
	if !ok {
		panic("type casting failed")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		lock, err := haBackend.LockWith(leadershipKey, "follow-the-leader")
		if err != nil {
			panic(err)
		}

		logger.Println("running for leadership")
		doneCh, err := lock.Lock(ctx.Done())
		if err != nil {
			panic(err)
		}
		logger.Println("elected as leader")

		select {
		case <-doneCh:
			logger.Println("lost leadership")
		case <-time.After(20 * time.Second):
			logger.Println("finished work, ceding leadership")
			if err := lock.Unlock(); err != nil {
				panic(err)
			}
		}
	}
}
