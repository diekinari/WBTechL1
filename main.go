package main

import (
	taskThree "WBTechL1/3"
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	taskThree.StartWorkers(3, ctx)

	select {
	case <-ctx.Done():
		fmt.Println("main goroutine has finished")
	}
}
