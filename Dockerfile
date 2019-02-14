# https://hub.docker.com/_/golang
ARG IMAGE_TAG="golang:1.11.5-stretch"
FROM ${IMAGE_TAG}

RUN apt-get update && apt-get install -y --no-install-recommends \
		zip \
		unzip \
	&& rm -rf /var/lib/apt/lists/*
WORKDIR "${GOPATH}/src/github.com/announce/altogether"
COPY . .
RUN make init
