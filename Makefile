generate:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        weather/weather.proto

_server:
	go run server/server.go

_client:
	go run client/client.go