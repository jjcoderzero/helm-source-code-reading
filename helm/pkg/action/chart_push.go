package action

import (
	"io"

	"helm.sh/helm/v3/internal/experimental/registry"
)

// ChartPush 执行Chart push操作.
type ChartPush struct {
	cfg *Configuration
}

// NewChartPush通过给出的配置创建一个新的ChartPush对象
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
