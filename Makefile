SECURE=false
ADDRESS=localhost:8080

build:
	go build -o bin/nb nb.go

protoc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	protoc --go_out=./internal/pkg  --go-grpc_out=./internal/pkg ./api/hexagon.proto

run-help:
	go run nb.go -h

run-help-repo:
	go run nb.go repo -h

run-help-hex:
	go run nb.go hex -h

run-repo-add:
	go run nb.go repo add 1000-0000-0000-0000 --secure=$(SECURE) --addr=$(ADDRESS)

run-repo-del:
	go run nb.go repo del 1000-0000-0000-0000 --secure=$(SECURE) --addr=$(ADDRESS)

run-hex-place:
	go run nb.go hex place 0 0 0 0000-0000-0000-0000  --secure=$(SECURE) --addr=$(ADDRESS)

run-hex-get:
	go run nb.go hex get 0 0 0 --radius=1 --secure=$(SECURE) --addr=$(ADDRESS)

run-hex-rm:
	go run nb.go 0 0 0 --secure=$(SECURE) --addr=$(ADDRESS)

run-hex-info:
	go run nb.go hex info 0 0 0 --secure=$(SECURE) --addr=$(ADDRESS)

run-status-server:
	go run nb.go status server --secure=$(SECURE) --addr=$(ADDRESS)

run-status-storage:
	go run nb.go status storage --secure=$(SECURE) --addr=$(ADDRESS)

run-status-clients:
	go run nb.go status clients --secure=$(SECURE) --addr=$(ADDRESS)
