FROM golang:1.21-alpine as build

ENV \
    CGO_ENABLED=1

RUN apk add --no-cache --update ca-certificates git gcc musl-dev bash curl make file jq && \
    rm -rf /var/cache/apk/*

RUN go version && \
    mkdir -p "/build" && \
    go install github.com/swaggo/swag/cmd/swag@v1.16.2 && \
    swag --version && \
    curl -o /usr/local/bin/swagger -L'#' $(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url') && \
    chmod +x /usr/local/bin/swagger

WORKDIR /build
ADD . /build
ADD go.sum go.mod /build/
RUN go mod download && \
    git config --global --add safe.directory /build

RUN make build

FROM alpine:3.17.2
RUN apk add --no-cache --update ca-certificates && \
    rm -rf /var/cache/apk/*

COPY --from=build /build/bin/sfs /srv/sfs
