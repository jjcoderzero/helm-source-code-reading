package chartutil

import "github.com/Masterminds/semver/v3"

// IsCompatibleRange compares a version to a constraint.
// It returns true if the version matches the constraint, and false in all other cases.
func IsCompatibleRange(constraint, ver string) bool {
	sv, err := semver.NewVersion(ver)
	if err != nil {
		return false
	}

	c, err := semver.NewConstraint(constraint)
	if err != nil {
		return false
	}
	return c.Check(sv)
}
