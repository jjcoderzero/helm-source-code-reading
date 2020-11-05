package plugin // import "helm.sh/helm/v3/pkg/plugin"

// Types of hooks
const (
	// Install is executed after the plugin is added.
	Install = "install"
	// Delete is executed after the plugin is removed.
	Delete = "delete"
	// Update is executed after the plugin is updated.
	Update = "update"
)

// Hooks is a map of events to commands.
type Hooks map[string]string
