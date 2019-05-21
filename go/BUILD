load("@io_bazel_rules_go//go:def.bzl", "go_prefix", "gazelle")

go_prefix("github.com/iv-p/site-builder")

# bazel rule definition
gazelle(
  prefix = "github.com/iv-p/site-builder",
  name = "gazelle",
  command = "fix",
)
