FROM golang:1.24-bullseye AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y \
    build-essential \
    pkg-config \
    librdkafka-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main .

FROM golang:1.24-alpine AS test-runner

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

FROM debian:bullseye-slim AS runtime

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
