LDFLAGS := -s -w
OS=$(shell uname -s)

ifneq ($(OS),Darwin)
	LDFLAGS += -extldflags "-static"
endif

.PHONY: build
build:
	mkdir -p bin
	export CGO_ENABLED=1
	go build -ldflags='${LDFLAGS}' -o bin/sfs mkuznets.com/go/sfs/cmd/sfs

.PHONY: generate
generate:
	./scripts/make-generate.sh

.PHONY: fmt
fmt:
	./scripts/make-fmt.sh

.PHONY: tidy
tidy:
	go mod tidy
	(cd api && go mod tidy)

.PHONY: run
run: build
	mkdir -p data
	bin/sfs run

.PHONY: test
test:
	go test -v ./...

.PHONY: distclean
distclean:
	rm -rf bin data
