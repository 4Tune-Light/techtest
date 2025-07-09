# Build stage
FROM golang:1.24.4 AS builder

WORKDIR /app

# Copy dependency files first
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project
COPY . .

# Build main application
RUN CGO_ENABLED=0 GOOS=linux go build -o techtest cmd/main.go

# Build migration tool (explicit path)
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/migrate-bin ./migrate

# Run stage
FROM alpine:latest

WORKDIR /root/

# Copy binaries
COPY --from=builder /app/techtest .
COPY --from=builder /app/migrate-bin ./migrate

# Copy supporting files
COPY --from=builder /app/internal ./internal
COPY --from=builder /app/.env .

# Install runtime dependencies
RUN apk add --no-cache make postgresql-client

EXPOSE 8080