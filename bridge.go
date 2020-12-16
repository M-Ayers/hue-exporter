package main

import (
	"fmt"
	"net/url"
	"path"
)

type Bridge struct {
	Url    string
	ApiKey string
}

func New(url, apiKey string) *Bridge {
	return &Bridge{url, apiKey}
}

func (b *Bridge) GetApiEndpoint(str ...string) (string, error) {
	b.Url = fmt.Sprintf("%s%s", "http://", b.Url)
	builtUrl, err := url.Parse(b.Url)
	if err != nil {
		return "URL Parsing error", err
	}

	builtUrl.Path = path.Join(builtUrl.Path, "/api/", b.ApiKey)

	for _, pathPart := range str {
		builtUrl.Path = path.Join(builtUrl.Path, pathPart)
	}
	return builtUrl.String(), nil
}
