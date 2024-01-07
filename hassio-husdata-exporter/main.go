package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"husdata_exporter/config"
	"husdata_exporter/metrics"
	"husdata_exporter/models"

	vm "github.com/VictoriaMetrics/metrics"
)

func main() {
	// Initalize config
	conf := config.New()

	// Set LogLevel
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: conf.LogLevel})
	slog.SetDefault(slog.New(h))

	// Initialize VPData
	vpdata := new(models.VPData)
	vpdataMutex := &sync.Mutex{} // Mutex for thread-safe access to vpdata

	// Initialize an HTTP client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Poll data immediately on startup
	pollData(client, conf, vpdata, vpdataMutex)

	// Start ticker and periodic polling
	ticker := time.NewTicker(time.Duration(conf.PollInterval) * time.Second)
	go func() {
		for range ticker.C {
			pollData(client, conf, vpdata, vpdataMutex)
		}
	}()

	// Initialize Mux
	mux := http.NewServeMux()

	// Expose vpdata metrics at `/` path.
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		vm.WritePrometheus(w, true)
	})

	// Expose the vpdata registered metrics at `/metrics` path.
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, req *http.Request) {
		vm.WritePrometheus(w, true)
	})

	// Define server
	server := http.Server{
		Addr:    ":8099",
		Handler: mux,
	}

	// Define extServer
	extServer := http.Server{
		Addr:    ":9101",
		Handler: mux,
	}

	// Start server
	go func() {
		slog.Info("Starting ingress http server", "port", 8099)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			slog.Error("HTTP server error", "error", err)
			os.Exit(1)
		}
		slog.Info("Stopped serving new ingress connections")
	}()

	if conf.ExternalServer {
		// Start ext http server
		go func() {
			slog.Info("Starting ext http server", "port", conf.Port)
			if err := extServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				slog.Error("Ext HTTP server error", "error", err)
				os.Exit(1)
			}
			slog.Info("Stopped serving new connections on ext.")
		}()
	}

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	slog.Info("Initializing graceful shutdown")
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("HTTP shutdown error", "error", err)
		os.Exit(1)
	}

	if conf.ExternalServer {
		if err := extServer.Shutdown(shutdownCtx); err != nil {
			slog.Error("HTTP shutdown error", "error", err)
			os.Exit(1)
		}
	}

	slog.Info("Graceful shutdown complete")
}

func pollData(client *http.Client, conf *config.Config, vpdata *models.VPData, vpdataMutex *sync.Mutex) {
	requestURL := fmt.Sprintf("http://%s/api/alldata", conf.HusdataHostname)
	slog.Debug("Polling husdata", "url", requestURL)

	res, err := client.Get(requestURL)
	if err != nil {
		slog.Error("error making http request", "error", err)
		return
	}

	if res.StatusCode != http.StatusOK {
		slog.Error("non-OK HTTP status", "status", res.StatusCode)
		res.Body.Close()
		return
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		slog.Error("error reading response body", "error", err)
		return
	}

	vpdataMutex.Lock()
	defer vpdataMutex.Unlock()
	if err := json.Unmarshal(resBody, vpdata); err != nil {
		slog.Error("error unmarshalling response body", "error", err)
	} else {
		metrics.Update(vpdata, conf.Lang)
	}
}
