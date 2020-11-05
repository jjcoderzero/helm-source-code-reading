package urlutil

import (
	"net/url"
	"path"
	"path/filepath"
)

// URLJoin joins a base URL to one or more path components.
//
// It's like filepath.Join for URLs. If the baseURL is pathish, this will still
// perform a join.
//
// If the URL is unparsable, this returns an error.
func URLJoin(baseURL string, paths ...string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	// We want path instead of filepath because path always uses /.
	all := []string{u.Path}
	all = append(all, paths...)
	u.Path = path.Join(all...)
	return u.String(), nil
}

// Equal normalizes two URLs and then compares for equality.
func Equal(a, b string) bool {
	au, err := url.Parse(a)
	if err != nil {
		a = filepath.Clean(a)
		b = filepath.Clean(b)
		// If urls are paths, return true only if they are an exact match
		return a == b
	}
	bu, err := url.Parse(b)
	if err != nil {
		return false
	}

	for _, u := range []*url.URL{au, bu} {
		if u.Path == "" {
			u.Path = "/"
		}
		u.Path = filepath.Clean(u.Path)
	}
	return au.String() == bu.String()
}

// ExtractHostname returns hostname from URL
func ExtractHostname(addr string) (string, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return "", err
	}
	return u.Hostname(), nil
}
