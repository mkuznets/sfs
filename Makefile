LDFLAGS := "-s -w"

build: sps

sps: swagger
	export CGO_ENABLED=0
	mkdir -p bin
	go build -ldflags=${LDFLAGS} -o bin/sps mkuznets.com/go/sps/cmd/sps

swagger:
	swag init -g internal/api/api.go
	swagger generate client --spec docs/swagger.json --name sps --strict-responders --target ./api

fmt:
	go fmt ./...
	swag fmt --exclude internal/api/resources.go

server: sps
	bin/sps server

db-up: sps
	bin/sps db up

db-down: sps
	bin/sps db down

db-status: sps
	bin/sps db status

test:
	go test -v ./...

.PHONY: sps build swagger
