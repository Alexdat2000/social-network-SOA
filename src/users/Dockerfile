FROM golang:1.24 AS build
WORKDIR /src/users
COPY go ./go
COPY tests ./tests
COPY main.go .
COPY go.* .

ENV CGO_ENABLED=0
RUN go get -v ./...

RUN go build -a -installsuffix cgo -o swagger .

FROM scratch AS runtime
COPY --from=build /src/users/swagger ./
ENTRYPOINT ["./swagger"]

FROM build AS test
CMD ["go", "test", "-v", "./tests/..."]
