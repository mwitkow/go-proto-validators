# Copyright 2016 Michal Witkowski. All Rights Reserved.
# See LICENSE for licensing terms.

export PATH := ${GOPATH}/bin:${PATH}

install:
	@echo "Installing govalidators to GOPATH"
	go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators

regenerate_test:
	@echo "Regenerating test .proto files"
	(protoc  \
	--proto_path=${GOPATH}/src \
	--proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
 	-proto_path=. \
	--gogo_out=. \
	--govalidators_out=. test/*.proto)

regenerate_example: install
	@echo "Regenerating example directory"
	(protoc  \
	--proto_path=${GOPATH}/src \
	--proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
	--proto_path=. \
	--go_out=. \
	--govalidators_out=. examples/*.proto)

test: install regenerate_test
	@echo "Running tests"
	(go test -v ./...)

regenerate:
	@echo "Regenerating validator.proto"
	(protoc \
	--proto_path=${GOPATH}/src \
	--proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
	--proto_path=. \
	--gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. \
	validator.proto)
