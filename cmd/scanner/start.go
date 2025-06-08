package scanner

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/raphael-d-cordeiro/reprocessing-engine/internal/application/scanner"
	"github.com/raphael-d-cordeiro/reprocessing-engine/internal/infra/cache"
	"github.com/raphael-d-cordeiro/reprocessing-engine/internal/infra/httputil"
	"github.com/raphael-d-cordeiro/reprocessing-engine/internal/infra/order"
)

func StartScanner(ctx context.Context) {
	// Initialize the scanner service
	// This is where you would set up your scanner logic, e.g., scanning files, processing data, etc.
	httpClient := httputil.New("http://localhost:8080")
	cacheClient, err := cache.New("localhost", 6379, "")
	if err != nil {
		log.Fatalf("Failed to connect to cache: %v", err)
	}
	orderRepository := order.New(httpClient)

	worker := scanner.New(cacheClient, orderRepository)
	worker.Start(ctx)

}

func Run(ctx context.Context) {
	ctx, cancelFunc := context.WithCancel(ctx)

	go StartScanner(ctx)
	HandleOSSignal(cancelFunc)
	log.Println("All services shutdown gracefully. Exiting... Scanner Service")

}

func HandleOSSignal(cancel context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	sig := <-signals
	signal.Stop(signals)
	log.Printf("Received signal: %s, Initiating graceful shutdown...\n", sig)
	cancel()
}
