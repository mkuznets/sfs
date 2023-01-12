LDFLAGS := "-s -w"

build: sps

sps: swagger
	export CGO_ENABLED=0
	mkdir -p bin
	go build -ldflags=${LDFLAGS} -o bin/sps mkuznets.com/go/sps/cmd/sps

swagger:
	swag init -g internal/sps/api/router.go

fmt:
	swag fmt

server: sps
	bin/sps server

.PHONY: sps build swagger
