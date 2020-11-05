package monocular

import (
	"errors"
	"net/url"
)

// ErrHostnameNotProvided indicates the url is missing a hostname
var ErrHostnameNotProvided = errors.New("no hostname provided")

// Client represents a client capable of communicating with the Monocular API.
type Client struct {

	// The base URL for requests
	BaseURL string

	// The internal logger to use
	Log func(string, ...interface{})
}

// New creates a new client
func New(u string) (*Client, error) {

	// Validate we have a URL
	if err := validate(u); err != nil {
		return nil, err
	}

	return &Client{
		BaseURL: u,
		Log:     nopLogger,
	}, nil
}

var nopLogger = func(_ string, _ ...interface{}) {}

// Validate if the base URL for monocular is valid.
func validate(u string) error {

	// Check if it is parsable
	p, err := url.Parse(u)
	if err != nil {
		return err
	}

	// Check that a host is attached
	if p.Hostname() == "" {
		return ErrHostnameNotProvided
	}

	return nil
}
