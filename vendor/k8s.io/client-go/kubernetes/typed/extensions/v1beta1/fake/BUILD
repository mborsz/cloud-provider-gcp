load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "fake_daemonset.go",
        "fake_deployment.go",
        "fake_deployment_expansion.go",
        "fake_extensions_client.go",
        "fake_ingress.go",
        "fake_podsecuritypolicy.go",
        "fake_replicaset.go",
        "fake_scale.go",
        "fake_scale_expansion.go",
    ],
    importmap = "io_k8s_cloud_provider_gcp/vendor/k8s.io/client-go/kubernetes/typed/extensions/v1beta1/fake",
    importpath = "k8s.io/client-go/kubernetes/typed/extensions/v1beta1/fake",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/api/extensions/v1beta1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/labels:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/types:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/watch:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes/typed/extensions/v1beta1:go_default_library",
        "//vendor/k8s.io/client-go/rest:go_default_library",
        "//vendor/k8s.io/client-go/testing:go_default_library",
    ],
)
