FROM golang:1.24-alpine

RUN apk update && apk add --no-cache gcc musl-dev git

WORKDIR /src/stats
COPY go.* .
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 go build -tags musl -ldflags '-extldflags "-static"' .

ENTRYPOINT ["./stats"]
