package api_order

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func StartAPIOrder(ctx context.Context) {
	r := gin.Default()
	r.GET("/orders", handler.GetOrders)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to start server: %v\n", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
	log.Println("Server exited gracefully")

}

func Run(ctx context.Context) {
	ctx, cancelFunc := context.WithCancel(ctx)

	go StartAPIOrder(ctx)
	HandleOSSignal(cancelFunc)
	log.Println("All services shutdown gracefully. Exiting... API Order Service")

}

func HandleOSSignal(cancel context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	sig := <-signals
	signal.Stop(signals)
	log.Printf("Received signal: %s, Initiating graceful shutdown...\n", sig)
	cancel()
}
