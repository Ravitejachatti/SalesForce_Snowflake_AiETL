package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"

    "salesforce-etl-ai/config"
    "salesforce-etl-ai/etl"
    "salesforce-etl-ai/salesforce"
    "salesforce-etl-ai/snowflake"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("❌ Failed to load config: %v", err)
    }

    if cfg.EnableMetrics {
        go config.StartMetricsServer(cfg.Port)
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigCh
        log.Println("🔌 Gracefully shutting down...")
        cancel()
    }()

    sf, err := salesforce.NewClient(cfg)
    if err != nil {
        log.Fatalf("❌ Salesforce login error: %v", err)
    }

    db, err := snowflake.Connect(cfg)
    if err != nil {
        log.Fatalf("❌ Snowflake connection error: %v", err)
    }
    defer db.DB.Close()

    if err := etl.ListenAndServe(ctx, sf, db); err != nil {
        log.Fatalf("❌ ETL failed: %v", err)
    }
}
