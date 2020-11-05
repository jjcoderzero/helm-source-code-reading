package action

import (
	"io"

	"helm.sh/helm/v3/internal/experimental/registry"
	"helm.sh/helm/v3/pkg/chart"
)

// ChartSave performs a chart save operation.
type ChartSave struct {
	cfg *Configuration
}

// NewChartSave creates a new ChartSave object with the given configuration.
func NewChartSave(cfg *Configuration) *ChartSave {
	return &ChartSave{
		cfg: cfg,
	}
}

// Run executes the chart save operation
func (a *ChartSave) Run(out io.Writer, ch *chart.Chart, ref string) error {
	r, err := registry.ParseReference(ref)
	if err != nil {
		return err
	}

	// If no tag is present, use the chart version
	if r.Tag == "" {
		r.Tag = ch.Metadata.Version
	}

	return a.cfg.RegistryClient.SaveChart(ch, r)
}
