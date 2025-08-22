# Friday Bot - Telegram Bot for Weekly Image Posts

A Golang Telegram bot that automatically posts images every Friday using scheduled tasks.

## Features

- 🤖 Telegram Bot API integration
- 📅 Automatic Friday posting with configurable time
- 🖼️ Random image selection from local directory
- ⏰ Configurable scheduling system
- 💬 Interactive command interface
- 🛡️ Error handling and logging
- ⚙️ Environment-based configuration

## Quick Start

### Prerequisites

- Go 1.21 or higher
- Telegram Bot Token (from @BotFather)
- Chat ID where the bot should post images

### Installation

1. Clone or download the project files
2. Set up your environment variables:

```bash
export BOT_TOKEN="your_telegram_bot_token"
export CHAT_ID="your_chat_id"  # Optional: can be set in config.json
```

3. Add images to the `sample_images` directory (supports .jpg, .jpeg, .png, .gif, .webp, .svg)
4. Run the bot:

```bash
go mod tidy
go run main.go
```

### Bot Token Setup

1. Open Telegram and search for @BotFather
2. Send `/newbot` command
3. Follow the instructions to create your bot
4. Copy the bot token and set it as BOT_TOKEN environment variable
5. Add your bot to a chat or group and get the chat ID
6. Set CHAT_ID environment variable or update config.json

### Configuration

The bot can be configured through environment variables or `config.json`:

- `BOT_TOKEN`: Your Telegram bot token (required)
- `CHAT_ID`: Target chat ID for posting images
- `POST_HOUR`: Hour to post (0-23, default: 9)
- `POST_MINUTE`: Minute to post (0-59, default: 0)  
- `IMAGES_DIR`: Directory containing images (default: "sample_images")
- `CHECK_INTERVAL`: How often to check for posting time in minutes (default: 30)
- `TIMEZONE`: Timezone for scheduling (default: "UTC")

### Commands

The bot supports these interactive commands:

- `/start` - Welcome message and bot introduction
- `/help` - Show available commands
- `/status` - Check bot status and configuration
- `/test` - Post a test image immediately

### Sample Images

The repository includes sample SVG images in the `sample_images` directory. Add your own images in supported formats for the bot to randomly select from.
