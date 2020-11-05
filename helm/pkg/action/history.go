package action

import (
	"github.com/pkg/errors"

	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/release"
)

// History is the action for checking the release's ledger.
type History struct {
	cfg *Configuration

	Max     int
	Version int
}

// NewHistory用已给配置创建了一个新的History对象
func NewHistory(cfg *Configuration) *History {
	return &History{
		cfg: cfg,
	}
}

// Run executes 'helm history' against the given release.
func (h *History) Run(name string) ([]*release.Release, error) {
	if err := h.cfg.KubeClient.IsReachable(); err != nil { // 判断是否能连接到kubernetes cluster。
		return nil, err
	}

	if err := chartutil.ValidateReleaseName(name); err != nil { // 验证release名字
		return nil, errors.Errorf("release name is invalid: %s", name)
	}

	h.cfg.Log("getting history for release %s", name)
	return h.cfg.Releases.History(name)
}
