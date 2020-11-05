package action

import (
	"io"

	"helm.sh/helm/v3/internal/experimental/registry"
)

// ChartPush performs a chart push operation.
type ChartPush struct {
	cfg *Configuration
}

// NewChartPush creates a new ChartPush object with the given configuration.
func NewChartPush(cfg *Configuration) *ChartPush {
	return &ChartPush{
		cfg: cfg,
	}
}

// Run executes the chart push operation
func (a *ChartPush) Run(out io.Writer, ref string) error {
	r, err := registry.ParseReference(ref)
	if err != nil {
		return err
	}
	return a.cfg.RegistryClient.PushChart(r)
}
