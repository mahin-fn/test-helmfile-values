package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/mokhlesurr031/sre-project/backend/internal/connection"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"

	_resourceHttp "github.com/mokhlesurr031/sre-project/backend/modules/resource/delivery/http"
	_resourceRepository "github.com/mokhlesurr031/sre-project/backend/modules/resource/repository"
	_resourceUseCase "github.com/mokhlesurr031/sre-project/backend/modules/resource/usecase"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starting Server...",
	Long:  `Starting Server...`,
	Run:   server,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func server(cmd *cobra.Command, args []string) {
	log.Println("Running Application")

	if err := connection.ConnectDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	defer connection.CloseDB()

	srv := buildHTTP(cmd, args)

	// Create a channel to wait for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func(sr *http.Server) {
		if err := sr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}(srv)

	// Wait for OS signals before shutting down
	<-stop

	// Shutdown the server gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server stopped gracefully")

}

func buildHTTP(_ *cobra.Command, _ []string) *http.Server {
	r := chi.NewRouter()
	r.Get("/metrics", promhttp.Handler().ServeHTTP)

	db := connection.DefaultDB()
	_ = db

	resourceRepo := _resourceRepository.New(db)
	resourceUseCase := _resourceUseCase.New(resourceRepo)
	_resourceHttp.NewHTTPHandler(r, resourceUseCase)

	return &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", 8080),
		Handler: r,
	}
}
