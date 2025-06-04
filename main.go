// main.go serves as the entry point for the application
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/raphael-d-cordeiro/reprocessing-engine/cmd/api_order"
	"github.com/raphael-d-cordeiro/reprocessing-engine/cmd/consumer"
	"github.com/raphael-d-cordeiro/reprocessing-engine/cmd/scanner"
)

func main() {
	appType := flag.String("app", "", "Application type to run (scanner or consumer)")
	flag.Parse()

	if *appType == "" {
		fmt.Println("Error: must specify application type with -app flag")
		fmt.Println("Usage: ./github.com/raphael-d-cordeiro/reprocessing-engine -app=scanner|consumer")
		os.Exit(1)
	}

	ctx := context.Background()

	switch *appType {
	case "scanner":
		scanner.Run(ctx)
	case "consumer":
		consumer.Run(ctx)
	case "api_order":
		api_order.Run(ctx)
	default:
		fmt.Printf("Error: unknown application type '%s'\n", *appType)
		fmt.Println("Usage: ./github.com/raphael-d-cordeiro/reprocessing-engine -app=scanner|consumer")
		os.Exit(1)
	}
}
