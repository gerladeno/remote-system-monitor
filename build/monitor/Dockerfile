FROM golang:1.16-alpine3.13 as builder
ADD . /src/app
WORKDIR /src/app
RUN go mod download

ARG LDFLAGS
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "$LDFLAGS" -o monitor ./cmd/monitor

FROM ubuntu:18.04

COPY --from=builder /src/app/monitor /monitor

ENTRYPOINT ["./monitor"]
