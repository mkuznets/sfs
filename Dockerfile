FROM golang:1.24-alpine3.22 AS build

ENV \
    CGO_ENABLED=1

RUN --mount=type=cache,target=/var/cache/apk/ \
    apk add --no-cache --update ca-certificates git gcc musl-dev bash curl make file jq

RUN go version && \
    mkdir -p "/build" && \
    go install github.com/swaggo/swag/cmd/swag@v1.16.2 && \
    swag --version && \
    curl -o /usr/local/bin/swagger -L'#' $(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url') && \
    chmod +x /usr/local/bin/swagger

WORKDIR /src
ADD go.sum go.mod /src/

RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod download -x

ADD . /src/
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod/ \
    make build

FROM alpine:3.22
RUN apk add --no-cache --update ca-certificates && \
    rm -rf /var/cache/apk/*

COPY --from=build /src/bin/sfs /srv/sfs

CMD ["/srv/sfs", "run"]
