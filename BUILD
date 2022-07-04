package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_source")

go_library(
    name = "problems",
    srcs = ["problems.go"],
    importpath = "problems",
    deps = ["@io_bazel_rules_go//go/tools/bazel"],
)

go_library(
    name = "utils",
    srcs = ["utils.go"],
    importpath = "utils",
)

exports_files(["expectedAnswers.json"])

exports_files(glob(["inputs/**/*.txt"]))

YEARS = [
    2015,
    2016,
    2017,
    2018,
    2019,
    2020,
    2021,
]

go_binary(
    name = "advent",
    srcs = ["advent.go"],
    data = glob(["inputs/**/*.txt"]) + ["//year%d:probs" % y for y in YEARS],
    deps = [
        ":problems",
        ":utils",
    ],
)

go_source(
    name = "test_src",
    srcs = ["advent_test.go"],
)

test_suite(
    name = "tests",
    tests = ["//year%d:tests" % y for y in YEARS],
)
