LDFLAGS := "-s -w"

build: sps

sps: swagger
	mkdir -p bin
	export CGO_ENABLED=1
	go build -tags sqlite_omit_load_extension -ldflags=${LDFLAGS} -o bin/sps mkuznets.com/go/sps/cmd/sps

swagger:
	swag init -g internal/api/api.go --output internal/api/swagger
	swagger generate client --spec internal/api/swagger/swagger.json --name sps --strict-responders --target ./api

fmt:
	go fmt ./...
	swag fmt --exclude internal/api/resources.go

server: sps
	bin/sps server

test:
	go test -v ./...

.PHONY: sps build swagger web
