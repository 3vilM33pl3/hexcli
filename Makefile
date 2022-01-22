build:
	go build -o bin/nb nb.go

protoc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	protoc --go_out=./internal/pkg  --go-grpc_out=./internal/pkg ./api/hexagon.proto
