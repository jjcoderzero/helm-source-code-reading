package releaseutil // import "helm.sh/helm/v3/pkg/releaseutil"

import rspb "helm.sh/helm/v3/pkg/release"

// FilterFunc returns true if the release object satisfies
// the predicate of the underlying filter func.
type FilterFunc func(*rspb.Release) bool

// Check applies the FilterFunc to the release object.
func (fn FilterFunc) Check(rls *rspb.Release) bool {
	if rls == nil {
		return false
	}
	return fn(rls)
}

// Filter applies the filter(s) to the list of provided releases
// returning the list that satisfies the filtering predicate.
func (fn FilterFunc) Filter(rels []*rspb.Release) (rets []*rspb.Release) {
	for _, rel := range rels {
		if fn.Check(rel) {
			rets = append(rets, rel)
		}
	}
	return
}

// Any returns a FilterFunc that filters a list of releases
// determined by the predicate 'f0 || f1 || ... || fn'.
func Any(filters ...FilterFunc) FilterFunc {
	return func(rls *rspb.Release) bool {
		for _, filter := range filters {
			if filter(rls) {
				return true
			}
		}
		return false
	}
}

// All returns a FilterFunc that filters a list of releases
// determined by the predicate 'f0 && f1 && ... && fn'.
func All(filters ...FilterFunc) FilterFunc {
	return func(rls *rspb.Release) bool {
		for _, filter := range filters {
			if !filter(rls) {
				return false
			}
		}
		return true
	}
}

// StatusFilter filters a set of releases by status code.
func StatusFilter(status rspb.Status) FilterFunc {
	return FilterFunc(func(rls *rspb.Release) bool {
		if rls == nil {
			return true
		}
		return rls.Info.Status == status
	})
}
