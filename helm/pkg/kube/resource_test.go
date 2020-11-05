package kube // import "helm.sh/helm/v3/pkg/kube"

import (
	"testing"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/resource"
)

func TestResourceList(t *testing.T) {
	mapping := &meta.RESTMapping{
		Resource: schema.GroupVersionResource{Group: "group", Version: "version", Resource: "pod"},
	}

	info := func(name string) *resource.Info {
		return &resource.Info{Name: name, Mapping: mapping}
	}

	var r1, r2 ResourceList
	r1 = []*resource.Info{info("foo"), info("bar")}
	r2 = []*resource.Info{info("bar")}

	if r1.Get(info("bar")).Mapping.Resource.Resource != "pod" {
		t.Error("expected get pod")
	}

	diff := r1.Difference(r2)
	if len(diff) != 1 {
		t.Error("expected 1 result")
	}

	if !diff.Contains(info("foo")) {
		t.Error("expected diff to return foo")
	}

	inter := r1.Intersect(r2)
	if len(inter) != 1 {
		t.Error("expected 1 result")
	}

	if !inter.Contains(info("bar")) {
		t.Error("expected intersect to return bar")
	}
}
