package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetApiEndpoint(t *testing.T) {
	b := New("10.0.1.10", "abc123")
	endpoint, err := b.GetApiEndpoint("test")
	assert.Nil(t, err)
	assert.Equal(t, endpoint, "http://10.0.1.10/api/abc123/test")
}
