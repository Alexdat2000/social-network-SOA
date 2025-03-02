FROM golang:1.24 AS build
WORKDIR /src/users
COPY go ./go
COPY main.go .
COPY go.* .

ENV CGO_ENABLED=0
RUN go get -v ./...

RUN go build -a -installsuffix cgo -o swagger .

FROM scratch AS runtime
COPY --from=build /src/users/swagger ./
EXPOSE 8080/tcp
ENTRYPOINT ["./swagger"]
