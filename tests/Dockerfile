FROM golang:1.24-alpine

WORKDIR /tests
COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "test", "-v", "./..."]
