package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"helm.sh/helm/v3/pkg/provenance"
	"helm.sh/helm/v3/pkg/repo"
	"helm.sh/helm/v3/pkg/repo/repotest"
)

func TestDependencyBuildCmd(t *testing.T) {
	srv, err := repotest.NewTempServerWithCleanup(t, "testdata/testcharts/*.tgz")
	defer srv.Stop()
	if err != nil {
		t.Fatal(err)
	}

	rootDir := srv.Root()
	srv.LinkIndices()

	chartname := "depbuild"
	createTestingChart(t, rootDir, chartname, srv.URL())
	repoFile := filepath.Join(rootDir, "repositories.yaml")

	cmd := fmt.Sprintf("dependency build '%s' --repository-config %s --repository-cache %s", filepath.Join(rootDir, chartname), repoFile, rootDir)
	_, out, err := executeActionCommand(cmd)

	// In the first pass, we basically want the same results as an update.
	if err != nil {
		t.Logf("Output: %s", out)
		t.Fatal(err)
	}

	if !strings.Contains(out, `update from the "test" chart repository`) {
		t.Errorf("Repo did not get updated\n%s", out)
	}

	// Make sure the actual file got downloaded.
	expect := filepath.Join(rootDir, chartname, "charts/reqtest-0.1.0.tgz")
	if _, err := os.Stat(expect); err != nil {
		t.Fatal(err)
	}

	// In the second pass, we want to remove the chart's request dependency,
	// then see if it restores from the lock.
	lockfile := filepath.Join(rootDir, chartname, "Chart.lock")
	if _, err := os.Stat(lockfile); err != nil {
		t.Fatal(err)
	}
	if err := os.RemoveAll(expect); err != nil {
		t.Fatal(err)
	}

	_, out, err = executeActionCommand(cmd)
	if err != nil {
		t.Logf("Output: %s", out)
		t.Fatal(err)
	}

	// Now repeat the test that the dependency exists.
	if _, err := os.Stat(expect); err != nil {
		t.Fatal(err)
	}

	// Make sure that build is also fetching the correct version.
	hash, err := provenance.DigestFile(expect)
	if err != nil {
		t.Fatal(err)
	}

	i, err := repo.LoadIndexFile(filepath.Join(rootDir, "index.yaml"))
	if err != nil {
		t.Fatal(err)
	}

	reqver := i.Entries["reqtest"][0]
	if h := reqver.Digest; h != hash {
		t.Errorf("Failed hash match: expected %s, got %s", hash, h)
	}
	if v := reqver.Version; v != "0.1.0" {
		t.Errorf("mismatched versions. Expected %q, got %q", "0.1.0", v)
	}
}

func TestDependencyBuildCmdWithHelmV2Hash(t *testing.T) {
	chartName := "testdata/testcharts/issue-7233"

	cmd := fmt.Sprintf("dependency build '%s'", chartName)
	_, out, err := executeActionCommand(cmd)

	// Want to make sure the build can verify Helm v2 hash
	if err != nil {
		t.Logf("Output: %s", out)
		t.Fatal(err)
	}
}
