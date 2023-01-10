LDFLAGS := "-s -w"

build: sps

sps:
	export CGO_ENABLED=0
	go build -ldflags=${LDFLAGS} mkuznets.com/go/sps/cmd/sps

.PHONY: sps build
