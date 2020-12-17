package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSensorsURL(t *testing.T) {
	b := NewBridge("10.0.1.10", "abc123")
	endpoint := getSensorsURL(b)
	assert.Equal(t, endpoint, "http://10.0.1.10/api/abc123/sensors")
}

// func TestGetSensors(t *testing.T) {
// 	b := NewBridge("10.0.1.10", "abc123")
// 	result := getSensors(b)
// 	assert.NotNil(t, result)
// }
