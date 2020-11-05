package gates

import (
	"os"
	"testing"
)

const name string = "HELM_EXPERIMENTAL_FEATURE"

func TestIsEnabled(t *testing.T) {
	os.Unsetenv(name)
	g := Gate(name)

	if g.IsEnabled() {
		t.Errorf("feature gate shows as available, but the environment variable %s was not set", name)
	}

	os.Setenv(name, "1")

	if !g.IsEnabled() {
		t.Errorf("feature gate shows as disabled, but the environment variable %s was set", name)
	}
}

func TestError(t *testing.T) {
	os.Unsetenv(name)
	g := Gate(name)

	if g.Error().Error() != "this feature has been marked as experimental and is not enabled by default. Please set HELM_EXPERIMENTAL_FEATURE=1 in your environment to use this feature" {
		t.Errorf("incorrect error message. Received %s", g.Error().Error())
	}
}

func TestString(t *testing.T) {
	os.Unsetenv(name)
	g := Gate(name)

	if g.String() != "HELM_EXPERIMENTAL_FEATURE" {
		t.Errorf("incorrect string representation. Received %s", g.String())
	}
}
