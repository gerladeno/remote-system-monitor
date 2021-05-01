#!/bin/bash
for goos in linux windows darwin
do
    # shellcheck disable=SC2043
    for goarch in amd64
    do
        version="${goos}/${goarch}:$1"
        echo "${version}"
        GOOS=${goos} GOARCH=${goarch} go build -v -o ./bin/monitor_${goos}_${goarch} -ldflags "-X main.version=${version} -X main.goos=${goos}" ./cmd/monitor
        GOOS=${goos} GOARCH=${goarch} go build -v -o ./bin/client_${goos}_${goarch} ./cmd/client
    done
done