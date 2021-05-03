FROM golang:1.16-alpine3.13 as builder
ADD . /src/app
WORKDIR /src/app
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o client ./cmd/client

FROM alpine:3.13
COPY --from=builder /src/app/client /client

ENTRYPOINT ["./client"]