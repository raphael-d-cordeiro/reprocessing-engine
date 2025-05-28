FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Build binário para o aplicativo
RUN go build -o bin/main .

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/bin/main /app/main

ENV TZ=America/Sao_Paulo

CMD ["sh", "-c", "echo Use scanner ou consumer como entrypoint"]
