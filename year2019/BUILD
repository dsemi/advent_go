package(default_visibility = ["//visibility:public"])
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "intcode",
    importpath = "year2019/intcode",
    srcs = ["intcode.go"],
)

[go_binary(
    name = "year2019_%s" % d,
    linkmode = "plugin",
    srcs = [d],
    out = "year2019_%s.so" % d,
    deps = [
        ":intcode",
        "//:utils",
    ],
) for d in glob(["day*.go"])]

filegroup(
    name = "year2019",
    data = [":year2019_%s" % d for d in glob(["day*.go"])],
)
