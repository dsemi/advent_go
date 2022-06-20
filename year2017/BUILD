package(default_visibility = ["//visibility:public"])
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

[go_binary(
    name = "year2017_%s" % d,
    linkmode = "plugin",
    srcs = [d],
    out = "year2017_%s.so" % d,
    deps = ["//:utils"],
) for d in glob(["day*.go"])]

filegroup(
    name = "year2017",
    data = [":year2017_%s" % d for d in glob(["day*.go"])],
)
