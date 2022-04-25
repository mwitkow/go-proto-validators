# Aerostat Docker Images

ARG ALPINE_VERSION=3.15


FROM alpine:${ALPINE_VERSION} AS build-env

WORKDIR /build

# Remember to update './scripts/update-vendor.sh'.
ARG GO_VERSION=1.17.2
ARG PROTOC_VERSION=3.20.1
ARG PROTOC_GEN_GO_VERSION=1.27.1
ARG PROTOC_GEN_GOGO_VERSION=1.3.2

ARG TARGET_OS=linux
ARG TARGET_ARCH=amd64

ENV GOPATH=/go
ENV GOOS="${TARGET_OS}"
ENV GOARCH="${TARGET_ARCH}"
ENV GOFLAGS="-v -mod=readonly -mod=vendor"
ENV GO111MODULE=on
ENV CGO_ENABLED=0

ENV PATH="${GOPATH}/bin:/usr/local/bin:/usr/local/go/bin:${PATH}"

RUN set -eux; \
  apk add --no-cache \
    bash \
    git \
    jq \
    libc6-compat \
    make

RUN set -eux; \
  wget -O go.tar.gz "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz" ; \
  tar -C /usr/local/ -xzf go.tar.gz; \
  rm -f go.tar.gz

RUN set -eux; \
  wget -O protoc.zip "https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip"; \
  unzip -j protoc.zip; \
  mv protoc /usr/local/bin/; \
  rm -f protoc.tar.gz

RUN set -eux; \
  wget -O protoc-gen-go.tar.gz "https://github.com/protocolbuffers/protobuf-go/releases/download/v${PROTOC_GEN_GO_VERSION}/protoc-gen-go.v${PROTOC_GEN_GO_VERSION}.linux.amd64.tar.gz"; \
  tar -C /usr/local/bin/ -xzf protoc-gen-go.tar.gz; \
  rm -f protoc-gen-go.tar.gz

RUN GOFLAGS="-v" go install "github.com/gogo/protobuf/protoc-gen-gogo@v${PROTOC_GEN_GOGO_VERSION}"


FROM build-env AS build

COPY . ./

# Make needs to build static binaries.

RUN make regenerate_test_golang_nodep test_nodep
