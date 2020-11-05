package chartutil

import (
	"k8s.io/client-go/kubernetes/scheme"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	helmversion "helm.sh/helm/v3/internal/version"
)

var (
	// DefaultVersionSet is the default version set, which includes only Core V1 ("v1").
	DefaultVersionSet = allKnownVersions()

	// DefaultCapabilities is the default set of capabilities.
	DefaultCapabilities = &Capabilities{
		KubeVersion: KubeVersion{
			Version: "v1.18.0",
			Major:   "1",
			Minor:   "18",
		},
		APIVersions: DefaultVersionSet,
		HelmVersion: helmversion.Get(),
	}
)

// Capabilities describes the capabilities of the Kubernetes cluster.
type Capabilities struct {
	// KubeVersion is the Kubernetes version.
	KubeVersion KubeVersion
	// APIversions are supported Kubernetes API versions.
	APIVersions VersionSet
	// HelmVersion is the build information for this helm version
	HelmVersion helmversion.BuildInfo
}

// KubeVersion is the Kubernetes version.
type KubeVersion struct {
	Version string // Kubernetes version
	Major   string // Kubernetes major version
	Minor   string // Kubernetes minor version
}

// String implements fmt.Stringer
func (kv *KubeVersion) String() string { return kv.Version }

// GitVersion returns the Kubernetes version string.
//
// Deprecated: use KubeVersion.Version.
func (kv *KubeVersion) GitVersion() string { return kv.Version }

// VersionSet is a set of Kubernetes API versions.
type VersionSet []string

// Has returns true if the version string is in the set.
//
//	vs.Has("apps/v1")
func (v VersionSet) Has(apiVersion string) bool {
	for _, x := range v {
		if x == apiVersion {
			return true
		}
	}
	return false
}

func allKnownVersions() VersionSet {
	// We should register the built in extension APIs as well so CRDs are
	// supported in the default version set. This has caused problems with `helm
	// template` in the past, so let's be safe
	apiextensionsv1beta1.AddToScheme(scheme.Scheme)
	apiextensionsv1.AddToScheme(scheme.Scheme)

	groups := scheme.Scheme.PrioritizedVersionsAllGroups()
	vs := make(VersionSet, 0, len(groups))
	for _, gv := range groups {
		vs = append(vs, gv.String())
	}
	return vs
}
