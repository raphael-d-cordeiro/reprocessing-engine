// main.go serves as the entry point for the application
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"syscall"

	"reprocessing-engine/cmd/consumer"
	"reprocessing-engine/cmd/scheduler"
)

func main() {
	appType := flag.String("app", "", "Application type to run (scanner or consumer)")
	flag.Parse()

	if *appType == "" {
		fmt.Println("Error: must specify application type with -app flag")
		fmt.Println("Usage: ./reprocessing-engine -app=scanner|consumer")
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	switch *appType {
	case "scanner":
		scheduler.RunScanner(ctx)
	case "consumer":
		consumer.RunConsumer(ctx)
	default:
		fmt.Printf("Error: unknown application type '%s'\n", *appType)
		fmt.Println("Usage: ./reprocessing-engine -app=scanner|consumer")
		os.Exit(1)
	}
}
