package registry // import "helm.sh/helm/v3/internal/experimental/registry"

const (
	// HelmChartConfigMediaType is the reserved media type for the Helm chart manifest config
	HelmChartConfigMediaType = "application/vnd.cncf.helm.config.v1+json"

	// HelmChartContentLayerMediaType is the reserved media type for Helm chart package content
	HelmChartContentLayerMediaType = "application/tar+gzip"
)

// KnownMediaTypes returns a list of layer mediaTypes that the Helm client knows about
func KnownMediaTypes() []string {
	return []string{
		HelmChartConfigMediaType,
		HelmChartContentLayerMediaType,
	}
}
