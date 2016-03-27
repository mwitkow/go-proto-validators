# Copyright 2016 Michal Witkowski. All Rights Reserved.
# See LICENSE for licensing terms.

install:
	@echo "Installing govalidators to GOPATH"
	go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators

regenerate_test:
	@echo "Regenerating test .proto files"
	(PATH=${GOPATH}/bin:${PATH} protoc  \
	--proto_path=${GOPATH}/src:${GOPATH}/src/github.com/gogo/protobuf/protobuf/:. \
	--gogo_out=. \
	--govalidators_out=. test/*.proto)

test: install regenerate_test
	@echo "Running tests"
	(go test -v ./...)

regenerate:
	@echo "Regenerating validator.proto"
	(PATH=${GOPATH}/bin:${PATH}  protoc \
	--proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf/:. \
	--gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. \
	validator.proto)
