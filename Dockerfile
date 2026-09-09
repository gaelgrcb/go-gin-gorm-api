# Stage 1: Build
FROM golang:1.27.1-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/main ./cmd/api

# Stage 2: Runtime
FROM alpine:3.20

WORKDIR /app

# Run as non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/main /app/main

USER appuser

EXPOSE 8080

CMD ["./main"]