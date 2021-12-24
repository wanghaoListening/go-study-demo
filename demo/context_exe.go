package main

import (
	"context"
	"fmt"
	"time"
)

//如何优雅的控制子goroutine退出

func subRoutine(ctx context.Context) {

	for {
		fmt.Println("............")
		time.Sleep(time.Second)
		select {

		case <-ctx.Done():
			break

		default:

		}

	}
}

func mainRoutine() {

	ctx, cancel := context.WithCancel(context.Background())
	go subRoutine(ctx)
	time.Sleep(time.Second * 6)
	cancel()
	fmt.Println("")
}
