# Friday Bot - Deployment Instructions

## Quick Start

1. Download the appropriate binary for your system:
   - **Linux**: `friday-bot`
   - **Windows**: `friday-bot-windows.exe`
   - **macOS**: `friday-bot-macos`

2. Make it executable (Linux/macOS only):
   ```bash
   chmod +x friday-bot
   ```

3. Set up environment variables:
   ```bash
   export BOT_TOKEN="your_telegram_bot_token"
   export CHAT_ID="your_chat_id"
   ```

4. Create images directory and add your images:
   ```bash
   mkdir sample_images
   # Add your .jpg, .png, .gif, .webp, or .svg files here
   ```

5. Run the bot:
   ```bash
   ./friday-bot                    # Linux/macOS
   friday-bot-windows.exe          # Windows
   ```

## Configuration Options

The bot can be configured via environment variables or `config.json`:

### Environment Variables
- `BOT_TOKEN` - Telegram bot token (required)
- `CHAT_ID` - Target chat ID for posts
- `POST_HOUR` - Hour to post (0-23, default: 9)
- `POST_MINUTE` - Minute to post (0-59, default: 0)
- `IMAGES_DIR` - Images directory (default: "sample_images")
- `CHECK_INTERVAL` - Check interval in minutes (default: 30)
- `TIMEZONE` - Timezone (default: "UTC")

### config.json Example
```json
{
    "chat_id": -1001234567890,
    "post_hour": 9,
    "post_minute": 0,
    "images_dir": "my_images",
    "check_interval_minutes": 30,
    "timezone": "America/New_York"
}
```

## Running as a Service

### Linux (systemd)
Create `/etc/systemd/system/friday-bot.service`:
```ini
[Unit]
Description=Friday Bot
After=network.target

[Service]
Type=simple
User=your_user
WorkingDirectory=/path/to/friday-bot
Environment=BOT_TOKEN=your_token
Environment=CHAT_ID=your_chat_id
ExecStart=/path/to/friday-bot
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

Enable and start:
```bash
sudo systemctl enable friday-bot
sudo systemctl start friday-bot
```

### Windows (Task Scheduler)
1. Open Task Scheduler
2. Create Basic Task
3. Set trigger to "At startup"
4. Set action to start your `friday-bot-windows.exe`
5. Configure environment variables in the task properties

## Troubleshooting

- **Bot not responding**: Check BOT_TOKEN is correct
- **No images posted**: Ensure images directory exists and contains supported files
- **Wrong timezone**: Set TIMEZONE environment variable
- **Permission denied**: Make binary executable with `chmod +x`

## Commands Available
- `/start` - Welcome message
- `/help` - Show available commands
- `/status` - Check bot status
- `/test` - Post test image immediately