package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("//:generator.bzl", "generate_rules")

go_library(
    name = "intcode",
    srcs = ["intcode.go"],
    importpath = "year2019/intcode",
)

generate_rules(
    srcs = glob(["day*.go"]),
    year = "2019",
    deps = [
        ":intcode",
        "//:utils",
    ],
)
