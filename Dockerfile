# https://hub.docker.com/_/golang
FROM golang:1.11.5-stretch

WORKDIR ${GOPATH}/src/github.com/announce/altogether
COPY . .
RUN make init
