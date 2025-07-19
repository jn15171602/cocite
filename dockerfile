# Build stage
FROM golang:1.24.1 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o chat-forum ./cmd/server

# Run stage
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/chat-forum .
COPY internal/templates/ internal/templates/

EXPOSE 8080

CMD ["./cocite"]