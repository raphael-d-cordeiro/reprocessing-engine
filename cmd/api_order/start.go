package api_order

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/raphael-d-cordeiro/reprocessing-engine/internal/api_order/handler"

	_ "github.com/raphael-d-cordeiro/reprocessing-engine/internal/api_order/docs" // Import the generated docs package

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func StartAPIOrder(ctx context.Context) {
	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.GET("/orders", handler.GetOrders)
	}
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
