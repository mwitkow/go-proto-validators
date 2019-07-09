#!/usr/bin/env bash

function install_protobuf() {
    if [[ -z "${PROTOBUF_VERSION:-}" ]]; then
        echo "Please set the version of protobuf to use via the PROTOBUF_VERSION environment variable."
        exit 1
    fi

    local deps_dir os_string version
    version="${PROTOBUF_VERSION}"
    deps_dir="${PROJECT_DIR}/deps"
    case "$(uname -s | tr "[:upper:]" "[:lower:]")" in
        linux)
            os_string="linux"
            ;;
        darwin)
            os_string="osx"
            ;;
        *)
            echo "This platform is not supported for running this script."
            exit 1
            ;;
    esac

    if [[ ! -e "${deps_dir}/${version}.zip" ]]; then
        echo "Downloading and installing protoc ${version}."
        mkdir -p "${deps_dir}"

        pushd "${deps_dir}" || exit 1
        rm -rf "${version}.zip" "${version}.zip.tmp" "bin" "include"
        wget "https://github.com/google/protobuf/releases/download/v${version}/protoc-${version}-${os_string}-x86_64.zip" -O "${version}.zip.tmp"
        unzip -o "${version}.zip.tmp"
        chmod 755 "bin/protoc"
        mv "${version}.zip.tmp" "${version}.zip"
        popd || exit 1
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

    PATH="${GOBIN:-"${HOME}/go/bin"}:${PATH}"
    export PATH
}
