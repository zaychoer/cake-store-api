FROM golang:latest

ENV PROJECT_DIR=/cake-store \
  GO111MODULE=on \
  CGO_ENABLED=0

RUN apt update && apt install git

ARG VERSION="4.15.2"

RUN git clone --branch "v${VERSION}" --depth 1 --single-branch https://github.com/golang-migrate/migrate /tmp/go-migrate

WORKDIR /tmp/go-migrate

RUN set -x \
  && CGO_ENABLED=0 go build -tags 'mysql' -ldflags="-s -w" -o ./migrate ./cmd/migrate \
  && ./migrate -version

RUN cp /tmp/go-migrate/migrate /usr/bin/migrate

WORKDIR /cake-store
RUN mkdir "/build"
COPY . .
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon -build="go build -o /build/cake-store" -command="/build/cake-store"

