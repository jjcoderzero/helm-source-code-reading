package chart

// Maintainer描述一个Chart maintainer.
type Maintainer struct {
	Name string `json:"name,omitempty"` // Name是一个用户名或者组织名
	Email string `json:"email,omitempty"` // Email是一个可选的电子邮件地址，用于联系指定的维护者
	URL string `json:"url,omitempty"`// URL是指向指定维护者地址的可选URL
}

// Chart文件的Metadata。这为Chart.yaml文件的结构建模
type Metadata struct {
	Name string `json:"name,omitempty"` 	// chart的名字
	Home string `json:"home,omitempty"`	// 一个SemVer 2符合版本的Chart字符串
	Sources []string `json:"sources,omitempty"` 	// 一个SemVer 2符合版本的Chart字符串
	Version string `json:"version,omitempty"`	// 一个SemVer 2符合版本的Chart字符串
	Description string `json:"description,omitempty"` 	// 用一句话描述Chart
	Keywords []string `json:"keywords,omitempty"`// 字符串关键字的列表
	Maintainers []*Maintainer `json:"maintainers,omitempty"`	// 维护者的名称和URL/电子邮件地址组合列表
	Icon string `json:"icon,omitempty"`	// 图标文件的URL
	APIVersion string `json:"apiVersion,omitempty"` // 要检查以启用Chart的条件
	Condition string `json:"condition,omitempty"` // 要检查以启用Chart的条件
	Tags string `json:"tags,omitempty"` // 要检查以启用Chart的标签
	AppVersion string `json:"appVersion,omitempty"` // 应用程序的版本包含在这个Chart中。
	Deprecated bool `json:"deprecated,omitempty"` // 这个Chart是否被废弃
	Annotations map[string]string `json:"annotations,omitempty"` // 注释是Helm未解释的附加映射，可供其他应用程序检查。
	KubeVersion string `json:"kubeVersion,omitempty"` // KubeVersion是一个SemVer约束，指定所需的Kubernetes版本。
	Dependencies []*Dependency `json:"dependencies,omitempty"` // Dependencies是Chart依赖项的列表。
	Type string `json:"type,omitempty"` // 指定Chart类型:应用程序或库
}

// Validate检查元数据中已知的问题，如果元数据不正确，则返回错误
func (md *Metadata) Validate() error {
	if md == nil {
		return ValidationError("chart.metadata is required")
	}
	if md.APIVersion == "" {
		return ValidationError("chart.metadata.apiVersion is required")
	}
	if md.Name == "" {
		return ValidationError("chart.metadata.name is required")
	}
	if md.Version == "" {
		return ValidationError("chart.metadata.version is required")
	}
	if !isValidChartType(md.Type) {
		return ValidationError("chart.metadata.type must be application or library")
	}

	// 这里需要对别名进行验证，以确保别名不包含任何非法字符。
	for _, dependency := range md.Dependencies { // 循环dependencies结构体
		if err := validateDependency(dependency); err != nil {
			return err
		}
	}

	// TODO validate valid semver here?
	return nil
}

func isValidChartType(in string) bool {
	switch in {
	case "", "application", "library": // Chart类型只能是application，library或者为空
		return true
	}
	return false
}

// validateDependency检查Chart中依赖数据结构的常见问题。此检查必须在加载依赖关系Chart之前在加载时完成。
func validateDependency(dep *Dependency) error {
	if len(dep.Alias) > 0 && !aliasNameFormat.MatchString(dep.Alias) { // 检测alias的长度和格式是否正确
		return ValidationErrorf("dependency %q has disallowed characters in the alias", dep.Name)
	}
	return nil
}
