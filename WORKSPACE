workspace(name = "com_github_mwitkow_go_proto_validators")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

GO_VERSION="1.13.5"

http_archive(
    name = "rules_proto",
    sha256 = "296ffd3e7992bd83fa75151255f7c7f27d22d6e52e2fd3c3d3d10c292317fbed",
    strip_prefix = "rules_proto-f6c112fa4eb2b8f934feb938a6fce41425e41587",
    urls = [
        "https://github.com/bazelbuild/rules_proto/archive/f6c112fa4eb2b8f934feb938a6fce41425e41587.tar.gz",
    ],
)

http_archive(
    name = "com_google_protobuf",
    strip_prefix = "protobuf-3.10.1",
    url = "https://github.com/protocolbuffers/protobuf/archive/v3.10.1.zip",
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "d5de13e9a994527b6dc41f39ad9ceee3214974dacb18f73a5fa2a4458ae6d3c9",
    strip_prefix = "rules_go-0.20.3",
    url = "https://github.com/bazelbuild/rules_go/archive/v0.20.3.tar.gz",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains(go_version = GO_VERSION)

http_archive(
    name = "bazel_gazelle",
    strip_prefix = "bazel-gazelle-0.19.1",
    urls = [
        "https://github.com/bazelbuild/bazel-gazelle/archive/v0.19.1.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    version = "v1.3.1",
)
