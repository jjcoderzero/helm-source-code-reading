package action

import (
	"fmt"
	"io"
	"path/filepath"

	"helm.sh/helm/v3/internal/experimental/registry"
	"helm.sh/helm/v3/pkg/chartutil"
)

// ChartExport执行一个chart导出操作.
type ChartExport struct {
	cfg *Configuration

	Destination string
}

// NewChartExport 使用给定的配置创建一个新的ChartExport对象.
func NewChartExport(cfg *Configuration) *ChartExport {
	return &ChartExport{
		cfg: cfg,
	}
}

// Run 执行chart导出操作
func (a *ChartExport) Run(out io.Writer, ref string) error {
	r, err := registry.ParseReference(ref)
	if err != nil {
		return err
	}

	ch, err := a.cfg.RegistryClient.LoadChart(r)
	if err != nil {
		return err
	}

	// 将chart保存到本地目标目录
	err = chartutil.SaveDir(ch, a.Destination)
	if err != nil {
		return err
	}

	d := filepath.Join(a.Destination, ch.Metadata.Name)
	fmt.Fprintf(out, "Exported chart to %s/\n", d)
	return nil
}
