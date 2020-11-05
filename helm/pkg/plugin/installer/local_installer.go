package installer // import "helm.sh/helm/v3/pkg/plugin/installer"

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// LocalInstaller installs plugins from the filesystem.
type LocalInstaller struct {
	base
}

// NewLocalInstaller creates a new LocalInstaller.
func NewLocalInstaller(source string) (*LocalInstaller, error) {
	src, err := filepath.Abs(source)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get absolute path to plugin")
	}
	i := &LocalInstaller{
		base: newBase(src),
	}
	return i, nil
}

// Install creates a symlink to the plugin directory.
//
// Implements Installer.
func (i *LocalInstaller) Install() error {
	if !isPlugin(i.Source) {
		return ErrMissingMetadata
	}
	debug("symlinking %s to %s", i.Source, i.Path())
	return os.Symlink(i.Source, i.Path())
}

// Update updates a local repository
func (i *LocalInstaller) Update() error {
	debug("local repository is auto-updated")
	return nil
}
