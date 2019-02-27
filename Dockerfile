# https://hub.docker.com/_/golang
FROM golang:1.12.0-stretch

WORKDIR ${GOPATH}/src/github.com/announce/altogether
COPY . .
RUN make init
