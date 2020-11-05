package release

// UninstallReleaseResponse represents a successful response to an uninstall request.
type UninstallReleaseResponse struct {
	// Release is the release that was marked deleted.
	Release *Release `json:"release,omitempty"`
	// Info is an uninstall message
	Info string `json:"info,omitempty"`
}
