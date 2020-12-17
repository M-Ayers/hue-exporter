package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetApiEndpoint(t *testing.T) {
	b := NewBridge("10.0.1.10", "abc123")
	endpoint, err := b.GetAPIEndpoint("test")
	assert.Nil(t, err)
	assert.Equal(t, endpoint, "http://10.0.1.10/api/abc123/test")
}
