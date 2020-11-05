package main

import (
	"runtime"
	"testing"
)

func TestDependencyListCmd(t *testing.T) {
	noSuchChart := cmdTestCase{
		name:      "No such chart",
		cmd:       "dependency list /no/such/chart",
		golden:    "output/dependency-list-no-chart-linux.txt",
		wantError: true,
	}

	noDependencies := cmdTestCase{
		name:   "No dependencies",
		cmd:    "dependency list testdata/testcharts/alpine",
		golden: "output/dependency-list-no-requirements-linux.txt",
	}

	if runtime.GOOS == "windows" {
		noSuchChart.golden = "output/dependency-list-no-chart-windows.txt"
		noDependencies.golden = "output/dependency-list-no-requirements-windows.txt"
	}

	tests := []cmdTestCase{noSuchChart,
		noDependencies, {
			name:   "Dependencies in chart dir",
			cmd:    "dependency list testdata/testcharts/reqtest",
			golden: "output/dependency-list.txt",
		}, {
			name:   "Dependencies in chart archive",
			cmd:    "dependency list testdata/testcharts/reqtest-0.1.0.tgz",
			golden: "output/dependency-list-archive.txt",
		}}
	runTestCmd(t, tests)
}

func TestDependencyFileCompletion(t *testing.T) {
	checkFileCompletion(t, "dependency", false)
}
