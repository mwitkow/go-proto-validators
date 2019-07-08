#!/usr/bin/env bash

function install_protobuf() {
    local deps_dir version
    version="${PROTOBUF_VERSION:-"3.0.2"}"
    deps_dir="${PROJECT_DIR}/deps"

    if [[ ! -e "${deps_dir}/${version}.zip" ]]; then
        echo "Downloading and installing protoc ${version}."
        mkdir -p "${deps_dir}"

        pushd "${deps_dir}"
        rm -rf "${version}.zip" "${version}.zip.tmp" "bin" "include"
        wget "https://github.com/google/protobuf/releases/download/v${version}/protoc-${version}-linux-x86_64.zip" -O "${version}.zip.tmp"
        unzip -o "${version}.zip.tmp"
        chmod 755 "bin/protoc"
        mv "${version}.zip.tmp" "${version}.zip"
        popd
    else
        echo "Reusing existing protoc ${version}."
    fi
}

function setup_proto_deps() {
    local dep dep_dir dep_location proto_deps
    dep_dir="${PROJECT_DIR}/deps"

    proto_deps=(
        "github.com/gogo/protobuf"
        "github.com/golang/protobuf"
        "github.com/mwitkow/go-proto-validators"
    )

    # Set up the target directory for symlinking in dependencies.
    mkdir -p "${dep_dir}"

    # Ensure the module dependencies are available
    go mod download

    for dep in "${proto_deps[@]}"; do
        dep_location="$(go list -m -json "${dep}" | jq -rM '.Dir')"

        mkdir -p "${dep_dir}/$(dirname "${dep}")"
        rm -f "${dep_dir}/${dep}"
        ln -sf "${dep_location}" "${dep_dir}/${dep}"
    done

    go install github.com/gogo/protobuf/protoc-gen-gogo
    go install github.com/golang/protobuf/protoc-gen-go
    go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators

    PATH="${GOBIN:-"${HOME}/go/bin"}:${PATH}"
    export PATH
}
