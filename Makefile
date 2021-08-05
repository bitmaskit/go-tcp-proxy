all:
	go test ./... && make echo && make proxy

echo:
	go build -o bin/echoserver ./cmd/echo/server.go

proxy:
	go build -o bin/proxyserver ./cmd/proxy/server.go

test:
	go test -v ./...

clean:
	rm -rf ./bin

.PHONY: all echo proxy test clean
