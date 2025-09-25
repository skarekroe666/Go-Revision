package chapter7

import (
	"context"
	"fmt"
	"time"
)

func Context() {
	ctx, cancel := context.WithCancel(context.Background())
	go doWork(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
