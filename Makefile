LDFLAGS := -s -w
OS=$(shell uname -s)

ifneq ($(OS),Darwin)
	LDFLAGS += -extldflags "-static"
endif

build: sfs

sfs: swagger
	make fmt tidy
	mkdir -p bin
	export CGO_ENABLED=1
	go build -ldflags='${LDFLAGS}' -o bin/sfs mkuznets.com/go/sfs/cmd/sfs

swagger:
	swag init -g internal/api/api.go --output internal/api/swagger
	swagger generate client --template=stratoscale --spec internal/api/swagger/swagger.json --name SimpleFeedService --strict-responders --target ./api

fmt:
	go fmt ./...
	swag fmt --exclude internal/api/resources.go

tidy:
	go mod tidy
	(cd api && go mod tidy)

run: sfs
	mkdir -p data
	bin/sfs run

test:
	go test -v ./...

distclean:
	rm -rf bin data

precommit:
	make swagger fmt tidy test

.PHONY: sfs build swagger web distclean test tidy fmt precommit
