# Overview

Friday Bot is a Golang Telegram bot designed to automatically post images every Friday at a configured time. The bot serves as an automated content scheduler that selects random images from a local directory and posts them to specified Telegram chats. It features a configurable scheduling system with timezone support and interactive command interface for user interaction.

# User Preferences

Preferred communication style: Simple, everyday language.

# System Architecture

## Core Application Design
The bot follows a modular architecture built around Telegram's Bot API integration. The main application loop combines scheduled task execution with real-time command processing, allowing both automated posting and interactive user commands.

## Scheduling System
Uses a configurable time-based scheduler that checks at regular intervals (default 30 minutes) whether it's time to post. The scheduling logic supports timezone configuration and precise hour/minute control for posting times, making it adaptable to different geographical locations and user preferences.

## Image Management
Implements a local file-based image storage system where images are stored in a designated directory (`sample_images` by default). The bot randomly selects images from this directory for posting, providing variety in automated content without external dependencies.

## Configuration Management
Adopts a hybrid configuration approach combining environment variables for sensitive data (BOT_TOKEN) with JSON file configuration for operational settings (posting schedule, image directory, timezone). This separation ensures security while maintaining ease of configuration management.

## Error Handling and Logging
Incorporates comprehensive error handling throughout the application with logging capabilities to track bot operations, posting attempts, and potential failures. This ensures reliable operation and easier debugging.

# External Dependencies

## Telegram Bot API
Primary integration with Telegram's Bot API for message sending and command processing. Requires a bot token obtained from @BotFather and chat ID configuration for target channels or groups.

## Runtime Environment
Built for Go 1.21+ runtime environment, utilizing Go's standard library for file operations, time management, and HTTP communications with Telegram's API endpoints.

## File System
Depends on local file system access for image storage and configuration file management. No external database or cloud storage dependencies, keeping the architecture simple and self-contained.