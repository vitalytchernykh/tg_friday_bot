package main

import (
        "log"
        "os"
        "os/signal"
        "sync"
        "syscall"

        "friday-bot/bot"
        "friday-bot/config"
        "friday-bot/logger"
        "friday-bot/scheduler"
)

func main() {
        // Initialize logger
        logger.Init()
        logger.Info("Starting Friday Bot...")

        // Load configuration
        cfg, err := config.Load()
        if err != nil {
                log.Fatalf("Failed to load configuration: %v", err)
        }

        // Validate bot token
        if cfg.BotToken == "" {
                log.Fatal("BOT_TOKEN environment variable is required")
        }

        // Initialize bot
        telegramBot, err := bot.New(cfg)
        if err != nil {
                log.Fatalf("Failed to initialize bot: %v", err)
        }

        // Initialize scheduler
        fridayScheduler := scheduler.New(telegramBot, cfg)

        // Create wait group for graceful shutdown
        var wg sync.WaitGroup

        // Start bot in goroutine
        wg.Add(1)
        go func() {
                defer wg.Done()
                logger.Info("Starting Telegram bot...")
                if err := telegramBot.Start(); err != nil {
                        logger.Error("Bot error: %v", err)
                }
        }()

        // Start scheduler in goroutine
        wg.Add(1)
        go func() {
                defer wg.Done()
                logger.Info("Starting Friday scheduler...")
                fridayScheduler.Start()
        }()

        // Wait for interrupt signal
        c := make(chan os.Signal, 1)
        signal.Notify(c, os.Interrupt, syscall.SIGTERM)

        logger.Info("Friday Bot is running. Press Ctrl+C to exit.")
        <-c

        // Graceful shutdown
        logger.Info("Shutting down Friday Bot...")
        telegramBot.Stop()
        fridayScheduler.Stop()

        // Wait for all goroutines to finish
        wg.Wait()
        logger.Info("Friday Bot stopped successfully")
}
