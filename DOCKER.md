# Friday Bot - Docker Deployment

This guide explains how to run Friday Bot using Docker containers.

## Prerequisites

- Docker installed on your system
- Bot credentials (BOT_TOKEN and CHAT_ID)

## Quick Start with Docker

### 1. Download all project files to your local machine

### 2. Build the Docker image
```bash
docker build -t friday-bot .
```

### 2. Run the container
```bash
docker run -d \
  --name friday-bot \
  --restart unless-stopped \
  -e BOT_TOKEN="your_telegram_bot_token" \
  -e CHAT_ID="your_chat_id" \
  -v $(pwd)/sample_images:/home/friday/sample_images:ro \
  friday-bot
```

## Using Docker Compose (Recommended)

### 1. Set up environment variables
```bash
cp .env.example .env
# Edit .env with your bot credentials
```

### 2. Start the service
```bash
docker-compose up -d
```

### 3. View logs
```bash
docker-compose logs -f friday-bot
```

### 4. Stop the service
```bash
docker-compose down
```

## Configuration

### Environment Variables
- `BOT_TOKEN` - Your Telegram bot token (required)
- `CHAT_ID` - Target chat ID for posts (required)
- `POST_HOUR` - Hour to post (0-23, default: 9)
- `POST_MINUTE` - Minute to post (0-59, default: 0)
- `IMAGES_DIR` - Images directory (default: sample_images)
- `CHECK_INTERVAL` - Check interval in minutes (default: 30)
- `TIMEZONE` - Timezone (default: UTC)
- `DEBUG` - Enable debug logging (default: false)

### Custom Images
Place your images in the `sample_images` directory before building the container, or mount a custom directory:

```bash
docker run -d \
  --name friday-bot \
  -e BOT_TOKEN="your_token" \
  -e CHAT_ID="your_chat_id" \
  -v /path/to/your/images:/home/friday/sample_images:ro \
  friday-bot
```

### Timezone Configuration
Set the `TIMEZONE` environment variable:
```bash
docker run -d \
  --name friday-bot \
  -e BOT_TOKEN="your_token" \
  -e CHAT_ID="your_chat_id" \
  -e TIMEZONE="America/New_York" \
  friday-bot
```

## Docker Commands

### Build image
```bash
docker build -t friday-bot .
```

### Run container in background
```bash
docker run -d --name friday-bot friday-bot
```

### View logs
```bash
docker logs friday-bot
docker logs -f friday-bot  # Follow logs
```

### Stop and remove container
```bash
docker stop friday-bot
docker rm friday-bot
```

### Shell into running container
```bash
docker exec -it friday-bot sh
```

## Production Deployment

### Using Docker Compose with restart policy
```yaml
services:
  friday-bot:
    build: .
    restart: unless-stopped
    environment:
      - BOT_TOKEN=${BOT_TOKEN}
      - CHAT_ID=${CHAT_ID}
    volumes:
      - ./sample_images:/home/friday/sample_images:ro
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "3"
```

### Health Monitoring
The container includes a health check that monitors the bot process:
```bash
docker inspect --format='{{json .State.Health}}' friday-bot
```

## Troubleshooting

### Check container status
```bash
docker ps -a
```

### View container logs
```bash
docker logs friday-bot
```

### Debug container
```bash
docker run --rm -it friday-bot sh
```

### Rebuild after changes
```bash
docker-compose down
docker-compose build --no-cache
docker-compose up -d
```