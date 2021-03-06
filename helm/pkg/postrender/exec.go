package postrender

import (
	"bytes"
	"io"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
)

type execRender struct {
	binaryPath string
}

// NewExec returns a PostRenderer implementation that calls the provided binary.
// It returns an error if the binary cannot be found. If the path does not
// contain any separators, it will search in $PATH, otherwise it will resolve
// any relative paths to a fully qualified path
func NewExec(binaryPath string) (PostRenderer, error) {
	fullPath, err := getFullPath(binaryPath)
	if err != nil {
		return nil, err
	}
	return &execRender{fullPath}, nil
}

// Run the configured binary for the post render
func (p *execRender) Run(renderedManifests *bytes.Buffer) (*bytes.Buffer, error) {
	cmd := exec.Command(p.binaryPath)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	var postRendered = &bytes.Buffer{}
	var stderr = &bytes.Buffer{}
	cmd.Stdout = postRendered
	cmd.Stderr = stderr

	go func() {
		defer stdin.Close()
		io.Copy(stdin, renderedManifests)
	}()
	err = cmd.Run()
	if err != nil {
		return nil, errors.Wrapf(err, "error while running command %s. error output:\n%s", p.binaryPath, stderr.String())
	}

	return postRendered, nil
}

// getFullPath returns the full filepath to the binary to execute. If the path
// does not contain any separators, it will search in $PATH, otherwise it will
// resolve any relative paths to a fully qualified path
func getFullPath(binaryPath string) (string, error) {
	// NOTE(thomastaylor312): I am leaving this code commented out here. During
	// the implementation of post-render, it was brought up that if we are
	// relying on plguins, we should actually use the plugin system so it can
	// properly handle multiple OSs. This will be a feature add in the future,
	// so I left this code for reference. It can be deleted or reused once the
	// feature is implemented

	// Manually check the plugin dir first
	// if !strings.Contains(binaryPath, string(filepath.Separator)) {
	// 	// First check the plugin dir
	// 	pluginDir := helmpath.DataPath("plugins") // Default location
	// 	// If location for plugins is explicitly set, check there
	// 	if v, ok := os.LookupEnv("HELM_PLUGINS"); ok {
	// 		pluginDir = v
	// 	}
	// 	// The plugins variable can actually contain multple paths, so loop through those
	// 	for _, p := range filepath.SplitList(pluginDir) {
	// 		_, err := os.Stat(filepath.Join(p, binaryPath))
	// 		if err != nil && !os.IsNotExist(err) {
	// 			return "", err
	// 		} else if err == nil {
	// 			binaryPath = filepath.Join(p, binaryPath)
	// 			break
	// 		}
	// 	}
	// }

	// Now check for the binary using the given path or check if it exists in
	// the path and is executable
	checkedPath, err := exec.LookPath(binaryPath)
	if err != nil {
		return "", errors.Wrapf(err, "unable to find binary at %s", binaryPath)
	}

	return filepath.Abs(checkedPath)
}
