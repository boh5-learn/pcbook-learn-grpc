gen:
	protoc -I./proto proto/*.proto --go-grpc_out=pb --go_out=pb

clean:
	# On windows, switch into wsl before run make clean
	# Because `make` on windows use CMD, and CMD not support `rm` command
	rm pb/*.go

server:
	go run cmd/server/main.go -port 56002

client:
	go run cmd/client/main.go -address "0.0.0.0:56002"

test:
	go test -cover -race ./...

lint:
	golangci-lint run
