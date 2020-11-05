package action

import (
	"helm.sh/helm/v3/pkg/chartutil"
)

// GetValues is the action for checking a given release's values.
//
// It provides the implementation of 'helm get values'.
type GetValues struct {
	cfg *Configuration

	Version   int
	AllValues bool
}

// NewGetValues creates a new GetValues object with the given configuration.
func NewGetValues(cfg *Configuration) *GetValues {
	return &GetValues{
		cfg: cfg,
	}
}

// Run executes 'helm get values' against the given release.
func (g *GetValues) Run(name string) (map[string]interface{}, error) {
	if err := g.cfg.KubeClient.IsReachable(); err != nil {
		return nil, err
	}

	rel, err := g.cfg.releaseContent(name, g.Version)
	if err != nil {
		return nil, err
	}

	// If the user wants all values, compute the values and return.
	if g.AllValues {
		cfg, err := chartutil.CoalesceValues(rel.Chart, rel.Config)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	}
	return rel.Config, nil
}
