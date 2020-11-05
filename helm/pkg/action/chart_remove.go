package action

import (
	"io"

	"helm.sh/helm/v3/internal/experimental/registry"
)

// ChartRemove执行移除Chart的操作
type ChartRemove struct {
	cfg *Configuration
}

// 使用给定的配置创建一个新的ChartRemove对象.
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
