workspace(name = "com_github_mwitkow_go_proto_validators")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

GO_VERSION = "1.13.5"

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
    sha256 = "678d91d8a939a1ef9cb268e1f20c14cd55e40361dc397bb5881e4e1e532679b1",
    strip_prefix = "protobuf-3.10.1",
    url = "https://github.com/protocolbuffers/protobuf/archive/v3.10.1.zip",
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

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
    sha256 = "d987004a72697334a095bbaa18d615804a28280201a50ed6c234c40ccc41e493",
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
    sum = "h1:G8O7TerXerS4F6sx9OV7/nRfJdnXgHZu/S/7F2SN+UE=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:ZDRjVQ15GmhC3fiQ8ni8+OwkZQO4DARzQgrnXU1Liz8=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:6nsPYzhq5kReh6QImI3k5qWzO4PEbvbIW2cwSfR/6xs=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    sum = "h1:reN85Pxc5larApoH1keMBiu2GWtPqXQ1nc9gx+jOU+E=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:4G4v2dO3VZwixGIRoQ5Lfboy6nUhCyYzaqnIAPPhYs4=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:TivCn/peBQ7UY8ooIcPgZFpTNSz0Q2U6UrFlUfqbe0Q=",
    version = "v1.3.0",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:NIou6eNFigscvKJmsbyez16S2cIS6idossORlFtSt2E=",
    version = "v0.0.0-20181030221726-6c7e314b6563",
)
