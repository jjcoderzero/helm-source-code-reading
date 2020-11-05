package kube

import (
	"io"
	"time"

	v1 "k8s.io/api/core/v1"
)

// Interface 表示能够与Kubernetes API通信的客户端.KubernetesClient必须是并发安全的.
type Interface interface {
	// Create创建一个或多个资源。
	Create(resources ResourceList) (*Result, error)

	Wait(resources ResourceList, timeout time.Duration) error

	// Delete删除一个或多个资源
	Delete(resources ResourceList) (*Result, []error)

	// Watch the resource in reader until it is "ready". This method
	//
	// For Jobs, "ready" means the Job ran to completion (exited without error).
	// For Pods, "ready" means the Pod phase is marked "succeeded".
	// For all other kinds, it means the kind was created or modified without
	// error.
	WatchUntilReady(resources ResourceList, timeout time.Duration) error

	// Update更新一个或多个资源或者如果不存在就创建资源
	Update(original, target ResourceList, force bool) (*Result, error)

	// Build creates a resource list from a Reader
	//
	// reader must contain a YAML stream (one or more YAML documents separated
	// by "\n---\n")
	//
	// Validates against OpenAPI schema if validate is true.
	Build(reader io.Reader, validate bool) (ResourceList, error)

	// WaitAndGetCompletedPodPhase等待超时，直到pod进入完成阶段并返回所述阶段(PodSucceed或PodFailed qualified)。
	WaitAndGetCompletedPodPhase(name string, timeout time.Duration) (v1.PodPhase, error)

	// isReachable检测client是否能够连接到集群
	IsReachable() error
}

var _ Interface = (*Client)(nil)
