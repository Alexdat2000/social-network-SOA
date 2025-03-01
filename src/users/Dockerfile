FROM golang:1.24-alpine

WORKDIR /src/users
COPY go.* .
RUN go mod download
COPY * .
RUN go build -o users

ENTRYPOINT ["./users"]
