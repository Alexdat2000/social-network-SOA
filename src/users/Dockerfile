FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM golang:1.24-alpine AS test-runner

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

FROM alpine:3.17 AS runtime

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
