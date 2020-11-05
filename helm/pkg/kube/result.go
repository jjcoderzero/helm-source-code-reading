package kube

// Result contains the information of created, updated, and deleted resources
// for various kube API calls along with helper methods for using those
// resources
type Result struct {
	Created ResourceList
	Updated ResourceList
	Deleted ResourceList
}

// If needed, we can add methods to the Result type for things like diffing
