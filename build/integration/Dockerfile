FROM golang:1.16-alpine3.13 as builder

ADD . /src/app
WORKDIR /src/app
RUN go mod download

RUN GOOS=linux CGO_ENABLED=0 go test -c ./tests/integration/ -o integration-tests

FROM alpine:3.13
COPY --from=builder /src/app/integration-tests /integration-tests
ENTRYPOINT ["./integration-tests"]