# https://hub.docker.com/_/golang
ARG IMAGE_TAG="golang:1.11.5-stretch"
FROM ${IMAGE_TAG}

WORKDIR "${GOPATH}/src/github.com/announce/altogether"
COPY . .
RUN make init
