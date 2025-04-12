# Dockerfile

FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o 1337b04rd ./cmd/1337b04rd

# ─── Final image ────────────────────────────────
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/1337b04rd .

EXPOSE 8080

CMD ["./1337b04rd"]
