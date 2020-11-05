package installer // import "helm.sh/helm/v3/pkg/plugin/installer"

import (
	"path/filepath"

	"helm.sh/helm/v3/pkg/helmpath"
)

type base struct {
	Source string // Source是插件的引用
}

func newBase(source string) base {
	return base{source}
}

// Path 是插件将被安装的位置.
func (b *base) Path() string {
	if b.Source == "" {
		return ""
	}
	return helmpath.DataPath("plugins", filepath.Base(b.Source))
}
