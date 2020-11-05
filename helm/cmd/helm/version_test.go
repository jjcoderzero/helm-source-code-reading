package main

import (
	"testing"
)

func TestVersion(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "default",
		cmd:    "version",
		golden: "output/version.txt",
	}, {
		name:   "short",
		cmd:    "version --short",
		golden: "output/version-short.txt",
	}, {
		name:   "template",
		cmd:    "version --template='Version: {{.Version}}'",
		golden: "output/version-template.txt",
	}, {
		name:   "client",
		cmd:    "version --client",
		golden: "output/version-client.txt",
	}, {
		name:   "client shorthand",
		cmd:    "version -c",
		golden: "output/version-client-shorthand.txt",
	}}
	runTestCmd(t, tests)
}

func TestVersionFileCompletion(t *testing.T) {
	checkFileCompletion(t, "version", false)
}
