package gates

import (
	"fmt"
	"os"
)

// Gate是feature gate的名称.
type Gate string

// String返回feature gate的字符串表示形式。
func (g Gate) String() string {
	return string(g)
}

// IsEnabled确定是否启用某个feature gate.
func (g Gate) IsEnabled() bool {
	return os.Getenv(string(g)) != ""
}

func (g Gate) Error() error {
	return fmt.Errorf("this feature has been marked as experimental and is not enabled by default. Please set %s=1 in your environment to use this feature", g.String())
}
