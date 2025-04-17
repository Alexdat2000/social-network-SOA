FROM golang:1.24-alpine AS build

RUN apk update && apk add --no-cache gcc musl-dev git

WORKDIR /src/users
COPY go ./go
COPY tests ./tests
COPY main.go .
COPY go.* .

RUN go get -v ./...

RUN CGO_ENABLED=1 go build -tags musl -ldflags '-extldflags "-static"' -a -installsuffix cgo -o api .

FROM scratch AS runtime
COPY --from=build /src/users/api ./
ENTRYPOINT ["./api"]

FROM build AS test
CMD ["go", "test", "-v", "./tests/..."]
