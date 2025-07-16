# Dockerfile para el backend (Go)
FROM golang:1.24.4-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o server ./cmd/server/main/Main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY ./config ./config
COPY ./database ./database
EXPOSE 8080
CMD ["./server"]
