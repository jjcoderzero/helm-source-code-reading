package action

import (
	"io"
)

// ChartList 执行Chart列表操作.
type ChartList struct {
	cfg *Configuration
}

// NewChartList用给出的配置创建一个新的ChartList对象。
func NewChartList(cfg *Configuration) *ChartList {
	return &ChartList{
		cfg: cfg,
	}
}

// Run executes the chart list operation
func (a *ChartList) Run(out io.Writer) error {
	return a.cfg.RegistryClient.PrintChartTable()
}
