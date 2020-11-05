// Package postrender contains an interface that can be implemented for custom
// post-renderers and an exec implementation that can be used for arbitrary
// binaries and scripts
package postrender

import "bytes"

type PostRenderer interface {
	// Run expects a single buffer filled with Helm rendered manifests. It
	// expects the modified results to be returned on a separate buffer or an
	// error if there was an issue or failure while running the post render step
	Run(renderedManifests *bytes.Buffer) (modifiedManifests *bytes.Buffer, err error)
}
