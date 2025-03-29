FROM golang:1.24-alpine

WORKDIR /src/content
COPY go.* .
RUN go mod download
COPY . .
RUN go build .

ENTRYPOINT ["./content"]
