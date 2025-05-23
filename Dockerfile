# syntax=docker/dockerfile:1
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Build binários separados para scanner e consumer
RUN go build -o bin/scanner ./cmd/scanner
RUN go build -o bin/consumer ./cmd/consumer

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/bin/scanner /app/scanner
COPY --from=builder /app/bin/consumer /app/consumer

ENV TZ=America/Sao_Paulo

CMD ["sh", "-c", "echo Use scanner ou consumer como entrypoint"]
