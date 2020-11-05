package registry // import "helm.sh/helm/v3/internal/experimental/registry"

import (
	"github.com/deislabs/oras/pkg/auth"
)

type (
	// Authorizer 处理注册表认证操作
	Authorizer struct {
		auth.Client
	}
)
