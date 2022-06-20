package(default_visibility = ["//visibility:public"])
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "problems",
    importpath = "problems",
    srcs = ["problems.go"],
)

go_library(
    name = "utils",
    importpath = "utils",
    srcs = ["utils.go"],
)

BINARY_DATA_DEPS = glob(["inputs/**/*.txt"]) + [
    "//year2015",
    "//year2016",
    "//year2017",
    "//year2018",
    "//year2019",
    "//year2020",
    "//year2021",
]

go_binary(
    name = "advent",
    srcs = ["advent.go"],
    data = BINARY_DATA_DEPS,
    deps = [
        ":problems",
        ":utils",
    ],
)

go_test(
    name = "advent_test",
    srcs = ["advent_test.go"],
    data = BINARY_DATA_DEPS + ["expectedAnswers.json"],
    deps = [":problems"],
)
