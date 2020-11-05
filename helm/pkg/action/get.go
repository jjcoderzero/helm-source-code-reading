package action

import (
	"helm.sh/helm/v3/pkg/release"
)

// Get is the action for checking a given release's information.
//
// It provides the implementation of 'helm get' and its respective subcommands (except `helm get values`).
type Get struct {
	cfg *Configuration

	// Initializing Version to 0 will get the latest revision of the release.
	Version int
}

// NewGet creates a new Get object with the given configuration.
func NewGet(cfg *Configuration) *Get {
	return &Get{
		cfg: cfg,
	}
}

// Run executes 'helm get' against the given release.
func (g *Get) Run(name string) (*release.Release, error) {
	if err := g.cfg.KubeClient.IsReachable(); err != nil {
		return nil, err
	}

	return g.cfg.releaseContent(name, g.Version)
}
