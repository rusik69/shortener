# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy source code
COPY . .

# Download dependencies
RUN make get

# Build the application
RUN make build

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/bin/shortener .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./shortener"]
