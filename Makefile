version := "0.0.0"

build_all:
	./make.sh $(version)

gen:
	go generate ./api/

run:
	GOOS=linux go run ./cmd/monitor -p 3002

client1:
	GOOS=linux go run ./cmd/client -p 3002 -m 10 -n 2 -i "0001"

client2:
	GOOS=linux go run ./cmd/client -p 3002 -m 3 -n 3 -i "0002"

client3:
	GOOS=linux go run ./cmd/client -p 3002 -m 8 -n 8 -i "0003"

