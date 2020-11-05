package action

import (
	"io"

	"helm.sh/helm/v3/internal/experimental/registry"
)

// ChartRemove performs a chart remove operation.
type ChartRemove struct {
	cfg *Configuration
}

// NewChartRemove creates a new ChartRemove object with the given configuration.
func NewChartRemove(cfg *Configuration) *ChartRemove {
	return &ChartRemove{
		cfg: cfg,
	}
}

// Run executes the chart remove operation
func (a *ChartRemove) Run(out io.Writer, ref string) error {
	r, err := registry.ParseReference(ref)
	if err != nil {
		return err
	}
	return a.cfg.RegistryClient.RemoveChart(r)
}
