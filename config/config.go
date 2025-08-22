package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Config holds all configuration for the bot
type Config struct {
	BotToken      string `json:"bot_token"`
	ChatID        int64  `json:"chat_id"`
	PostHour      int    `json:"post_hour"`
	PostMinute    int    `json:"post_minute"`
	ImagesDir     string `json:"images_dir"`
	CheckInterval int    `json:"check_interval_minutes"`
	TimeZone      string `json:"timezone"`
}

// Load configuration from environment variables and config file
func Load() (*Config, error) {
	config := &Config{
		PostHour:      9,  // Default to 9 AM
		PostMinute:    0,  // Default to 0 minutes
		ImagesDir:     "sample_images",
		CheckInterval: 30, // Check every 30 minutes
		TimeZone:      "UTC",
	}

	// Load from config file if exists
	if err := loadFromFile(config); err != nil {
		// Config file is optional, just log the error
		fmt.Printf("Config file not found or invalid, using defaults: %v\n", err)
	}

	// Override with environment variables
	if token := os.Getenv("BOT_TOKEN"); token != "" {
		config.BotToken = token
	}

	if chatID := os.Getenv("CHAT_ID"); chatID != "" {
		if id, err := strconv.ParseInt(chatID, 10, 64); err == nil {
			config.ChatID = id
		}
	}

	if hour := os.Getenv("POST_HOUR"); hour != "" {
		if h, err := strconv.Atoi(hour); err == nil && h >= 0 && h <= 23 {
			config.PostHour = h
		}
	}

	if minute := os.Getenv("POST_MINUTE"); minute != "" {
		if m, err := strconv.Atoi(minute); err == nil && m >= 0 && m <= 59 {
			config.PostMinute = m
		}
	}

	if imagesDir := os.Getenv("IMAGES_DIR"); imagesDir != "" {
		config.ImagesDir = imagesDir
	}

	if interval := os.Getenv("CHECK_INTERVAL"); interval != "" {
		if i, err := strconv.Atoi(interval); err == nil && i > 0 {
			config.CheckInterval = i
		}
	}

	if tz := os.Getenv("TIMEZONE"); tz != "" {
		config.TimeZone = tz
	}

	return config, nil
}

// loadFromFile loads configuration from config.json file
func loadFromFile(config *Config) error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(config)
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.BotToken == "" {
		return fmt.Errorf("bot token is required")
	}

	if c.PostHour < 0 || c.PostHour > 23 {
		return fmt.Errorf("post hour must be between 0 and 23")
	}

	if c.PostMinute < 0 || c.PostMinute > 59 {
		return fmt.Errorf("post minute must be between 0 and 59")
	}

	if c.CheckInterval <= 0 {
		return fmt.Errorf("check interval must be positive")
	}

	return nil
}
