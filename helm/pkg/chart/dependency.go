package chart

import "time"

// Dependency描述另一个Chart所依赖的一个Chart。可以用来表达开发人员的意图，或者用来捕获Chart的状态。
type Dependency struct {
	Name string `json:"name"` // Name是依赖项的名称。这必须与依赖关系的Chart中的名称匹配。
	Version string `json:"version,omitempty"` // Version就是这个Chart的版本(范围)。lock文件将始终生成一个版本，而依赖项可能包含一个语义版本范围。
	Repository string `json:"repository"`// 到存储库的URL。将index.yaml附加到该字符串后，将得到一个可用于获取存储库索引的URL。
	Condition string `json:"condition,omitempty"` // 一个解析为布尔值的yaml路径，用于启用/禁用图表(例如subchart1.enabled)
	Tags []string `json:"tags,omitempty"` // Tags可以用来将Chart的启用/禁用组合在一起
	Enabled bool `json:"enabled,omitempty"` // Enabled bool决定是否应该加载Chart
	ImportValues []interface{} `json:"import-values,omitempty"` // ImportValues保存要导入的源值到父键的映射。每个项可以是一个字符串或一对子/父子列表项。
	Alias string `json:"alias,omitempty"` // Alias用于Chart的可用别名
}

// Lock是依赖项的锁文件。它表示依赖项应该处于的状态。
type Lock struct {
	Generated time.Time `json:"generated"` // Generated是锁文件最后生成的日期。
	Digest string `json:"digest"` // Digest是Chart.yaml中依赖项的hash。
	Dependencies []*Dependency `json:"dependencies"` // Dependencies是此锁定文件已锁定的依赖项的列表。
}
