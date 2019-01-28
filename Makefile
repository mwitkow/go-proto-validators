# Copyright 2016 Michal Witkowski. All Rights Reserved.
# See LICENSE for licensing terms.

export PATH := ${GOPATH}/bin:${PATH}

init:
	go get -u "github.com/gogo/protobuf/protoc-gen-gogo"

install:
	@echo "--- Installing govalidators to GOPATH"
	go install github.com/TheThingsIndustries/go-proto-validators/protoc-gen-govalidators


regenerate_test_generic: install
	@echo "Regenerating the generic test .proto files (with gogo)"
	(protoc  \
	--proto_path=${GOPATH}/src \
 	--proto_path=test \
	--gogo_out=test \
	--govalidators_out=gogoimport=true:test test/generic/*.proto)

regenerate_test_gogo: install
	@echo "Regenerating test .proto files with gogo imports"
	(protoc  \
	--proto_path=${GOPATH}/src \
 	--proto_path=test \
	--gogo_out=test/gogo \
	--govalidators_out=gogoimport=true:test/gogo test/*.proto)

regenerate_example: install
	@echo "--- Regenerating example directory"
	(protoc  \
	--proto_path=${GOPATH}/src \
	--proto_path=. \
	--go_out=. \
	--govalidators_out=. examples/*.proto)

errors: install
	@echo "--- Regenerating error definitions"
	(protoc  \
	--proto_path=${GOPATH}/src \
	--proto_path=errors \
	--go_out=errors errors/*.proto)

test: regenerate_test_generic regenerate_test_gogo
	@echo "Running tests"
	go test -v ./...

regenerate:
	@echo "--- Regenerating validator.proto"
	(protoc \
	--proto_path=${GOPATH}/src \
	--proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
	--proto_path=. \
	--gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. \
	validator.proto)
