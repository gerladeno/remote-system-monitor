version := "0.0.2"
LDFLAGS := -X main.version=$(version)

gen:
	go generate ./api/

build_all: gen
	./make.sh $(version)

run: gen
	GOOS=linux go run ./cmd/monitor -p 3002

client1:
	GOOS=linux go run ./cmd/client -p 3002 -m 10 -n 2 -i "0001"

client2:
	GOOS=linux go run ./cmd/client -p 3002 -m 3 -n 3 -i "0002"

client3:
	GOOS=linux go run ./cmd/client -p 3002 -m 8 -n 8 -i "0003"

docker_build: gen
	docker-compose build

docker_run:
	docker-compose up