# Build stage
FROM golang:1.19-alpine AS builder

# Install git for dependency resolution
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o friday-bot .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests to Telegram API
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user
RUN adduser -D -s /bin/sh friday

# Set working directory
WORKDIR /home/friday

# Copy binary from builder stage
COPY --from=builder /app/friday-bot .

# Copy configuration files
COPY config.json .
COPY sample_images/ ./sample_images/

# Create images directory and set permissions
RUN mkdir -p sample_images && \
    chown -R friday:friday /home/friday

# Switch to non-root user
USER friday

# Expose port (not actually used by this bot, but good practice)
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD pgrep friday-bot || exit 1

# Run the binary
CMD ["./friday-bot"]