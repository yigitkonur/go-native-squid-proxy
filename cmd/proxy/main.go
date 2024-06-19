package main

import (
    "log" // Standard library log package
    "net/http"
    "os"
    "os/signal"
    "syscall"

    "proxy-server/pkg/config"
    logger "proxy-server/pkg/log" // Alias to avoid redeclaration with standard log package
    "proxy-server/pkg/proxy"

    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Setup logging
    loggerInstance, err := logger.NewLogger(cfg.LogLevel)
    if err != nil {
        log.Fatalf("Error setting up logger: %v", err)
    }
    defer loggerInstance.Sync()
    sugar := loggerInstance.Sugar()

    // Setup metrics
    http.Handle("/metrics", promhttp.Handler())

    // Initialize and start the proxy server
    proxyServer := proxy.NewProxyServer(cfg, sugar)
    sugar.Infof("Starting proxy server on %s", cfg.ServerAddress)

    // Channel to handle OS signals for graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        if err := proxyServer.Start(); err != nil {
            sugar.Fatalf("Failed to start proxy server: %v", err)
        }
    }()

    // Wait for interrupt signal to gracefully shut down the server
    <-quit
    sugar.Info("Shutting down proxy server...")
    if err := proxyServer.Shutdown(); err != nil {
        sugar.Fatalf("Failed to gracefully shut down proxy server: %v", err)
    }
    sugar.Info("Proxy server stopped")
}
