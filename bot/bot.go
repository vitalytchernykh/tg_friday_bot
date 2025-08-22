package bot

import (
        "fmt"
        "friday-bot/config"
        "friday-bot/images"
        "friday-bot/logger"
        "strings"

        tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Bot represents the Telegram bot
type Bot struct {
        api           *tgbotapi.BotAPI
        config        *config.Config
        imageManager  *images.Manager
        updateChannel tgbotapi.UpdatesChannel
        stopChannel   chan struct{}
}

// New creates a new Bot instance
func New(cfg *config.Config) (*Bot, error) {
        if err := cfg.Validate(); err != nil {
                return nil, fmt.Errorf("invalid configuration: %w", err)
        }

        api, err := tgbotapi.NewBotAPI(cfg.BotToken)
        if err != nil {
                return nil, fmt.Errorf("failed to create bot API: %w", err)
        }

        logger.Info("Authorized on account %s", api.Self.UserName)

        imageManager := images.New(cfg.ImagesDir)

        return &Bot{
                api:          api,
                config:       cfg,
                imageManager: imageManager,
                stopChannel:  make(chan struct{}),
        }, nil
}

// Start begins listening for updates
func (b *Bot) Start() error {
        u := tgbotapi.NewUpdate(0)
        u.Timeout = 60

        updates := b.api.GetUpdatesChan(u)
        b.updateChannel = updates

        for {
                select {
                case update := <-updates:
                        b.handleUpdate(update)
                case <-b.stopChannel:
                        logger.Info("Bot stopped")
                        return nil
                }
        }
}

// Stop gracefully stops the bot
func (b *Bot) Stop() {
        close(b.stopChannel)
        b.api.StopReceivingUpdates()
}

// handleUpdate processes incoming updates
func (b *Bot) handleUpdate(update tgbotapi.Update) {
        if update.Message == nil {
                return
        }

        message := update.Message
        logger.Info("Received message from %s: %s", message.From.UserName, message.Text)

        // Handle commands
        if message.IsCommand() {
                b.handleCommand(message)
                return
        }

        // Handle regular messages
        b.handleMessage(message)
}

// handleCommand processes bot commands
func (b *Bot) handleCommand(message *tgbotapi.Message) {
        command := strings.ToLower(message.Command())

        switch command {
        case "start":
                b.sendMessage(message.Chat.ID, "🤖 Welcome to Friday Bot!\n\nI automatically post images every Friday. Here are available commands:\n\n/status - Check bot status\n/test - Post a test image\n/help - Show this help message")
        
        case "help":
                b.sendMessage(message.Chat.ID, "📋 Available commands:\n\n/start - Welcome message\n/status - Check bot status\n/test - Post a test image\n/help - Show this help\n\nI automatically post images every Friday at the configured time!")
        
        case "status":
                status := fmt.Sprintf("✅ Bot is running!\n\n📅 Next Friday post scheduled\n🕘 Post time: %02d:%02d\n📁 Images directory: %s\n🔄 Check interval: %d minutes",
                        b.config.PostHour, b.config.PostMinute, b.config.ImagesDir, b.config.CheckInterval)
                b.sendMessage(message.Chat.ID, status)
        
        case "test":
                if err := b.PostFridayImage(message.Chat.ID); err != nil {
                        b.sendMessage(message.Chat.ID, fmt.Sprintf("❌ Failed to post test image: %v", err))
                } else {
                        b.sendMessage(message.Chat.ID, "✅ Test image posted successfully!")
                }
        
        default:
                b.sendMessage(message.Chat.ID, "❓ Unknown command. Use /help to see available commands.")
        }
}

// handleMessage processes regular messages
func (b *Bot) handleMessage(message *tgbotapi.Message) {
        // For now, just acknowledge the message
        if strings.Contains(strings.ToLower(message.Text), "friday") {
                b.sendMessage(message.Chat.ID, "🎉 Yes, I love Fridays too! I'll post something special every Friday!")
        }
}

// sendMessage sends a text message to a chat
func (b *Bot) sendMessage(chatID int64, text string) {
        msg := tgbotapi.NewMessage(chatID, text)
        if _, err := b.api.Send(msg); err != nil {
                logger.Error("Failed to send message: %v", err)
        }
}

// PostFridayImage posts a Friday image to the specified chat
func (b *Bot) PostFridayImage(chatID int64) error {
        // If no chat ID specified, use the configured one
        if chatID == 0 {
                chatID = b.config.ChatID
        }

        imagePath, err := b.imageManager.GetRandomImage()
        if err != nil {
                return fmt.Errorf("failed to get image: %w", err)
        }

        // Create photo message
        photo := tgbotapi.NewPhoto(chatID, tgbotapi.FilePath(imagePath))
        photo.Caption = "🎉 Happy Friday! Have a wonderful weekend! 🎉"

        // Send the photo
        if _, err := b.api.Send(photo); err != nil {
                return fmt.Errorf("failed to send photo: %w", err)
        }

        logger.Info("Successfully posted Friday image to chat %d", chatID)
        return nil
}

// GetBotInfo returns bot information
func (b *Bot) GetBotInfo() string {
        return fmt.Sprintf("Bot: @%s (ID: %d)", b.api.Self.UserName, b.api.Self.ID)
}
