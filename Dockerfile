# Aerostat Docker Images

ARG ALPINE_VERSION=3.15


FROM alpine:${ALPINE_VERSION} AS build-env

WORKDIR /build

# Remember to update './scripts/update-vendor.sh'.
ARG GO_VERSION=1.17.2

ARG PROTOC_VERSION=3.18.1
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
    "protoc=$PROTOC_VERSION-r1"

RUN set -eux; \
  wget -O go.tar.gz "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz" ; \
  tar -C /usr/local/ -xzf go.tar.gz; \
  rm -f go.tar.gz

RUN set -eux; \
  wget -O protoc-gen-go.tar.gz "https://github.com/protocolbuffers/protobuf-go/releases/download/v${PROTOC_GEN_GO_VERSION}/protoc-gen-go.v${PROTOC_GEN_GO_VERSION}.linux.amd64.tar.gz"; \
  tar -C /usr/local/bin/ -xzf protoc-gen-go.tar.gz; \
  rm -f protoc-gen-go.tar.gz

RUN GOFLAGS="-v" go install "github.com/gogo/protobuf/protoc-gen-gogo@v${PROTOC_GEN_GOGO_VERSION}"


FROM build-env AS build

COPY . ./

# Make needs to build static binaries.

# RUN protoc \
#   --go_out=test/golang \
#   --govalidators_out=test/golang \
#   --proto_path=test \
#   test/*.proto

# RUN protoc \
#   --proto_path=vendor \
#   --proto_path=. \
#   --gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. \
#   validator.proto

  # --proto_path=deps \
  # --proto_path=deps/include \
  # --proto_path=deps/github.com/gogo/protobuf/protobuf \

  # --go_out=. \
  # --go_opt=paths=source_relative \
  # --go-grpc_out=. \
  # --go-grpc_opt=paths=source_relative \
  # --govalidators_out=. \
  # --govalidators_opt=paths=source_relative \
  # --proto_path=. \


# RUN go test -v ./...
