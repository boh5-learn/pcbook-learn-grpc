gen:
	protoc -I./proto proto/*.proto --go-grpc_out=pb --go_out=pb

clean:
	# On windows, switch into wsl before run make clean
	# Because `make` on windows use CMD, and CMD not support `rm` command
	rm pb/*.go

run:
	go run main.go

test:
	go test -cover -race ./...
