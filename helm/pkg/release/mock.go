package release

import (
	"fmt"
	"math/rand"

	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/time"
)

// MockHookTemplate is the hook template used for all mock release objects.
var MockHookTemplate = `apiVersion: v1
kind: Job
metadata:
  annotations:
    "helm.sh/hook": pre-install
`

// MockManifest is the manifest used for all mock release objects.
var MockManifest = `apiVersion: v1
kind: Secret
metadata:
  name: fixture
`

// MockReleaseOptions allows for user-configurable options on mock release objects.
type MockReleaseOptions struct {
	Name      string
	Version   int
	Chart     *chart.Chart
	Status    Status
	Namespace string
}

// Mock creates a mock release object based on options set by MockReleaseOptions. This function should typically not be used outside of testing.
func Mock(opts *MockReleaseOptions) *Release {
	date := time.Unix(242085845, 0).UTC()

	name := opts.Name
	if name == "" {
		name = "testrelease-" + fmt.Sprint(rand.Intn(100))
	}

	version := 1
	if opts.Version != 0 {
		version = opts.Version
	}

	namespace := opts.Namespace
	if namespace == "" {
		namespace = "default"
	}

	ch := opts.Chart
	if opts.Chart == nil {
		ch = &chart.Chart{
			Metadata: &chart.Metadata{
				Name:       "foo",
				Version:    "0.1.0-beta.1",
				AppVersion: "1.0",
			},
			Templates: []*chart.File{
				{Name: "templates/foo.tpl", Data: []byte(MockManifest)},
			},
		}
	}

	scode := StatusDeployed
	if len(opts.Status) > 0 {
		scode = opts.Status
	}

	info := &Info{
		FirstDeployed: date,
		LastDeployed:  date,
		Status:        scode,
		Description:   "Release mock",
		Notes:         "Some mock release notes!",
	}

	return &Release{
		Name:      name,
		Info:      info,
		Chart:     ch,
		Config:    map[string]interface{}{"name": "value"},
		Version:   version,
		Namespace: namespace,
		Hooks: []*Hook{
			{
				Name:     "pre-install-hook",
				Kind:     "Job",
				Path:     "pre-install-hook.yaml",
				Manifest: MockHookTemplate,
				LastRun:  HookExecution{},
				Events:   []HookEvent{HookPreInstall},
			},
		},
		Manifest: MockManifest,
	}
}
