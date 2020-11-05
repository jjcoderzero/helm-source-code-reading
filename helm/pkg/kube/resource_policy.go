package kube // import "helm.sh/helm/v3/pkg/kube"

// ResourcePolicyAnno is the annotation name for a resource policy
const ResourcePolicyAnno = "helm.sh/resource-policy"

// KeepPolicy is the resource policy type for keep
//
// This resource policy type allows resources to skip being deleted
//   during an uninstallRelease action.
const KeepPolicy = "keep"
