package(default_visibility = ["//visibility:public"])
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

[go_binary(
    name = "year2021_%s" % d,
    linkmode = "plugin",
    srcs = [d],
    out = "year2021_%s.so" % d,
    deps = ["//:utils"],
) for d in glob(["day*.go"])]

filegroup(
    name = "year2021",
    data = [":year2021_%s" % d for d in glob(["day*.go"])],
)
