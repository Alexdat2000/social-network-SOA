FROM golang:1.24-alpine AS build

WORKDIR /src/gateway
COPY go.* .
RUN go mod download
COPY . .
RUN go build .

FROM scratch AS runtime
COPY --from=build /src/gateway ./
ENTRYPOINT ["./gateway"]
