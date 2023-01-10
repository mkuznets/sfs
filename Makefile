LDFLAGS := "-s -w"

build: sps

sps:
	export CGO_ENABLED=0
	mkdir -p bin
	go build -ldflags=${LDFLAGS} -o bin/sps mkuznets.com/go/sps/cmd/sps

.PHONY: sps build
