package registry // import "helm.sh/helm/v3/internal/experimental/registry"

import (
	"github.com/containerd/containerd/remotes"
)

type (
	// Resolver provides remotes based on a locator
	Resolver struct {
		remotes.Resolver
	}
)
