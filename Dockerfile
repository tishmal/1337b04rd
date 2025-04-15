# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install required packages
RUN apk add --no-cache git make

# Copy go module files first for better layer caching
COPY go.mod go.sum* ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o 1337b04rd ./cmd/1337b04rd

# Stage 2: Runtime image
FROM alpine:3.19

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy the binary and required files from the builder stage
COPY --from=builder /app/1337b04rd /app/

COPY web/static/templates/ /app/web/static/templates/

# COPY static/ /app/static/

# Create a non-root user
RUN adduser -D -g '' appuser
RUN chown -R appuser:appuser /app
USER appuser

# Expose the application port
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/app/1337b04rd"]
CMD ["--port", "8080"]