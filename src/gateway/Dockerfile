FROM golang:1.24-alpine

WORKDIR /src/gateway
COPY go.* .
RUN go mod download
COPY * .
RUN go build -o gateway

ENTRYPOINT ["./gateway"]
