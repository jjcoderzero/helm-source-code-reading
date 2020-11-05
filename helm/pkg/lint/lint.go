package lint // import "helm.sh/helm/v3/pkg/lint"

import (
	"path/filepath"

	"helm.sh/helm/v3/pkg/lint/rules"
	"helm.sh/helm/v3/pkg/lint/support"
)

// All runs all of the available linters on the given base directory.
func All(basedir string, values map[string]interface{}, namespace string, strict bool) support.Linter {
	// Using abs path to get directory context
	chartDir, _ := filepath.Abs(basedir)

	linter := support.Linter{ChartDir: chartDir}
	rules.Chartfile(&linter)
	rules.ValuesWithOverrides(&linter, values)
	rules.Templates(&linter, values, namespace, strict)
	rules.Dependencies(&linter)
	return linter
}
