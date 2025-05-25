FROM golang:1.24-bullseye

RUN apt-get update && apt-get install -y gcc librdkafka-dev git

WORKDIR /tests
COPY go.* .
RUN go mod download
COPY . .

CMD ["go", "test", "-v"]
