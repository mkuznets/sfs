LDFLAGS := "-s -w"

build: sfs

sfs: swagger
	mkdir -p bin
	export CGO_ENABLED=1
	go build -tags sqlite_omit_load_extension -ldflags=${LDFLAGS} -o bin/sfs mkuznets.com/go/sfs/cmd/sfs

swagger:
	swag init -g internal/api/api.go --output internal/api/swagger
	swagger generate client --spec internal/api/swagger/swagger.json --name SimpleFeedService --strict-responders --target ./api

fmt:
	go fmt ./...
	swag fmt --exclude internal/api/resources.go

run: sfs
	bin/sfs run

test:
	go test -v ./...

.PHONY: sfs build swagger web
