package main

import (
	"fmt"
	"net/url"
	"path"
)

// Bridge struct contains URL and APIKey
type Bridge struct {
	URL    string
	APIKey string
}

// NewBridge returns a Bridge struct
func NewBridge(url, apiKey string) *Bridge {
	return &Bridge{url, apiKey}
}

// GetAPIEndpoint returns a built URL from this Bridge
func (b *Bridge) GetAPIEndpoint(str ...string) (string, error) {
	b.URL = fmt.Sprintf("%s%s", "http://", b.URL)
	builtURL, err := url.Parse(b.URL)
	if err != nil {
		return "URL Parsing error", err
	}

	builtURL.Path = path.Join(builtURL.Path, "/api/", b.APIKey)

	for _, pathPart := range str {
		builtURL.Path = path.Join(builtURL.Path, pathPart)
	}
	return builtURL.String(), nil
}
