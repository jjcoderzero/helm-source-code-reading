package chartutil

import (
	"regexp"

	"github.com/pkg/errors"
)

// validName是用于资源名称的正则表达式.
// According to the Kubernetes help text, the regular expression it uses is:
//	[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*
// 这遵循上面的正则表达式(但需要完整的字符串匹配，而不是部分匹配)
var validName = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`)

var (
	// errMissingName表示没有提供release(名称)。
	errMissingName = errors.New("no name provided")

	// errInvalidName表示提供了无效的版本名称
	errInvalidName = errors.New("invalid release name, must match regex ^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])+$ and the length must not longer than 53")

	// errInvalidKubernetesName表示该名称不符合Kubernetes对元数据名称的限制。
	errInvalidKubernetesName = errors.New("invalid metadata name, must match regex ^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])+$ and the length must not longer than 253")
)

const (
	// maxNameLen是helm允许release名字的最大长度
	maxReleaseNameLen = 53
	// maxMetadataNameLen是Kubernetes允许任何名称的最大长度。
	maxMetadataNameLen = 253
)

// ValidateReleaseName执行检查Helm发行版名称的条目
//
// 为了让Helm允许一个名字，它必须低于一个特定的字符数(53)并且匹配一个reguar表达式。
func ValidateReleaseName(name string) error {
	// 保留这种情况是为了向后兼容
	if name == "" {
		return errMissingName

	}
	if len(name) > maxReleaseNameLen || !validName.MatchString(name) { // 限制字符长度和是否匹配命名要求
		return errInvalidName
	}
	return nil
}

// ValidateMetadataName验证Kubernetes元数据对象的名称字段。
func ValidateMetadataName(name string) error {
	if name == "" || len(name) > maxMetadataNameLen || !validName.MatchString(name) { // 空字符串、长度超过253个字符的字符串或与regexp不匹配的字符串将失败。
		return errInvalidKubernetesName
	}
	return nil
}
