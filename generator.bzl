load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_test")

def generate_rules(year, srcs, deps = []):
    # d = "day01.go"
    for d in srcs:
        go_binary(
            name = d[:-3],
            srcs = [d],
            out = "%s.so" % d[:-3],
            linkmode = "plugin",
            deps = deps,
        )

        go_test(
            name = "%s_test" % d[:-3],
            srcs = ["//:test_src"],
            args = [
                "--year=%s" % year,
                "--day=%d" % int(d[-5:-3], 10),
            ],
            data = [
                "//:expectedAnswers.json",
                "//:inputs/%s/input%s.txt" % (year, int(d[-5:-3], 10)),
                ":%s" % d[:-3],
            ],
            deps = [
                "//:problems",
                "@io_bazel_rules_go//go/tools/bazel",
            ],
        )

    native.filegroup(
        name = "probs",
        data = [":%s" % d[:-3] for d in srcs],
    )

    native.test_suite(
        name = "tests",
    )
