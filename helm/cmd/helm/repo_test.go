package main

import (
	"testing"
)

func TestRepoFileCompletion(t *testing.T) {
	checkFileCompletion(t, "repo", false)
}
