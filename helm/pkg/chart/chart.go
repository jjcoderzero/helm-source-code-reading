package chart

import (
	"path/filepath"
	"regexp"
	"strings"
)

const APIVersionV1 = "v1" // APIVersionV1是版本1的API版本号。

const APIVersionV2 = "v2" // APIVersionV2是版本2的API版本号。

var aliasNameFormat = regexp.MustCompile("^[a-zA-Z0-9_-]+$") // aliasNameFormat定义别名中合法的字符。

// Chart是一个helm包，包含元数据、默认配置、零个或多个可选参数化模板，以及零个或更多Chart(依赖关系)。
type Chart struct {
	Raw []*File `json:"-"` // Raw包含Chart存档中最初包含的文件的原始内容。这应该不使用，除非在特殊情况下，如“helm show values”，我们想要显示原始值，注释和所有。
	Metadata *Metadata `json:"metadata"` // 元数据是Chart文件的内容
	Lock *Lock `json:"lock"` // Lock是Chart.lock的内容。
	Templates []*File `json:"templates"` // chart的Templates
	Values map[string]interface{} `json:"values"` // Values是此Chart的默认配置
	Schema []byte `json:"schema"` // Schema是一个可选的JSON模式，用于在值上强加结构
	Files []*File `json:"files"` // Files是Chart归档中的杂项文件，如自述文件、许可文件等。

	parent       *Chart
	dependencies []*Chart
}

type CRD struct {
	Name string // Name是crd file的File.Name
	Filename string // FileName是obj文件的名称，包括(子)chartfullpath
	File *File // File是crd的obj文件
}

// SetDependencies替换Chart依赖项。
func (ch *Chart) SetDependencies(charts ...*Chart) {
	ch.dependencies = nil
	ch.AddDependency(charts...)
}

// Name返回chart的名字
func (ch *Chart) Name() string {
	if ch.Metadata == nil {
		return ""
	}
	return ch.Metadata.Name
}

// AddDependency决定Chart是否是subChart。
func (ch *Chart) AddDependency(charts ...*Chart) {
	for i, x := range charts {
		charts[i].parent = ch
		ch.dependencies = append(ch.dependencies, x)
	}
}

// Root 寻找root chart.
func (ch *Chart) Root() *Chart {
	if ch.IsRoot() {
		return ch
	}
	return ch.Parent().Root()
}

// Dependencies 是此Chart所依赖的Chart.
func (ch *Chart) Dependencies() []*Chart { return ch.dependencies }

// IsRoot确定该Chart是否是根Chart
func (ch *Chart) IsRoot() bool { return ch.parent == nil }

// Parent 返回子Chart的父Chart
func (ch *Chart) Parent() *Chart { return ch.parent }

// ChartPath以点表示法返回此图表的完整路径
func (ch *Chart) ChartPath() string {
	if !ch.IsRoot() {
		return ch.Parent().ChartPath() + "." + ch.Name()
	}
	return ch.Name()
}

// ChartFullPath返回此Chart的完整路径
func (ch *Chart) ChartFullPath() string {
	if !ch.IsRoot() {
		return ch.Parent().ChartFullPath() + "/charts/" + ch.Name()
	}
	return ch.Name()
}

// Validate 验证元数据.
func (ch *Chart) Validate() error {
	return ch.Metadata.Validate()
}

// AppVersion返回Chart的appversion.
func (ch *Chart) AppVersion() string {
	if ch.Metadata == nil {
		return ""
	}
	return ch.Metadata.AppVersion
}

// CRDs返回Helm Chart的“crds/”目录下的文件对象列表
// Deprecated: use CRDObjects()
func (ch *Chart) CRDs() []*File {
	files := []*File{}
	// Find all resources in the crds/ directory
	for _, f := range ch.Files {
		if strings.HasPrefix(f.Name, "crds/") && hasManifestExtension(f.Name) {
			files = append(files, f)
		}
	}
	// Get CRDs from dependencies, too.
	for _, dep := range ch.Dependencies() {
		files = append(files, dep.CRDs()...)
	}
	return files
}

// CRDObjects returns a list of CRD objects in the 'crds/' directory of a Helm chart & subcharts
func (ch *Chart) CRDObjects() []CRD {
	crds := []CRD{}
	// Find all resources in the crds/ directory
	for _, f := range ch.Files {
		if strings.HasPrefix(f.Name, "crds/") && hasManifestExtension(f.Name) {
			mycrd := CRD{Name: f.Name, Filename: filepath.Join(ch.ChartFullPath(), f.Name), File: f}
			crds = append(crds, mycrd)
		}
	}
	// Get CRDs from dependencies, too.
	for _, dep := range ch.Dependencies() {
		crds = append(crds, dep.CRDObjects()...)
	}
	return crds
}

func hasManifestExtension(fname string) bool {
	ext := filepath.Ext(fname)
	return strings.EqualFold(ext, ".yaml") || strings.EqualFold(ext, ".yml") || strings.EqualFold(ext, ".json")
}
