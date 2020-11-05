package registry // import "helm.sh/helm/v3/internal/experimental/registry"

import (
	"context"
	"fmt"
	"io"
	"time"

	orascontext "github.com/deislabs/oras/pkg/context"
	units "github.com/docker/go-units"
	"github.com/sirupsen/logrus"
)

// byteCountBinary produces a human-readable file size
func byteCountBinary(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

// shortDigest returns first 7 characters of a sha256 digest
func shortDigest(digest string) string {
	if len(digest) == 64 {
		return digest[:7]
	}
	return digest
}

// timeAgo returns a human-readable timestamp representing time that has passed
func timeAgo(t time.Time) string {
	return units.HumanDuration(time.Now().UTC().Sub(t))
}

// ctx retrieves a fresh context.
// disable verbose logging coming from ORAS (unless debug is enabled)
func ctx(out io.Writer, debug bool) context.Context {
	if !debug {
		return orascontext.Background()
	}
	ctx := orascontext.WithLoggerFromWriter(context.Background(), out)
	orascontext.GetLogger(ctx).Logger.SetLevel(logrus.DebugLevel)
	return ctx
}
