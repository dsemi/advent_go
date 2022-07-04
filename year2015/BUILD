package(default_visibility = ["//visibility:public"])

load("//:generator.bzl", "generate_rules")

generate_rules(
    srcs = glob(["day*.go"]),
    year = "2015",
    deps = ["//:utils"],
)
