package action

import (
	"io"

	"helm.sh/helm/v3/internal/experimental/registry"
)

// ChartPull performs a chart pull operation.
type ChartPull struct {
	cfg *Configuration
}

// NewChartPull creates a new ChartPull object with the given configuration.
func NewChartPull(cfg *Configuration) *ChartPull {
	return &ChartPull{
		cfg: cfg,
	}
}

// Run executes the chart pull operation
func (a *ChartPull) Run(out io.Writer, ref string) error {
	r, err := registry.ParseReference(ref)
	if err != nil {
		return err
	}
	return a.cfg.RegistryClient.PullChart(r)
}
