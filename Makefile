version := "0.0.0"

build_all:
	./make.sh $(version)

run:
	GOOS=linux go run ./cmd/monitor

gen:
	go generate ./api/