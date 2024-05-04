package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformPostBody_Normal(t *testing.T) {
	result, err := transformPostBody([]byte(`{"key": "value", "redirect_uri": "https://example.com/abc?def=ghi"}`), "example2.com")
	expected := []byte(`{"key":"value","redirect_uri":"https://example2.com/abc?def=ghi"}`)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestTransformPostBody_Empty(t *testing.T) {
	result, err := transformPostBody([]byte(`{}`), "example2.com")
	expected := []byte(`{}`)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestTransformPostBody_NormalWithoutRedirectUri(t *testing.T) {
	result, err := transformPostBody([]byte(`{"key": "value"}`), "example2.com")
	expected := []byte(`{"key":"value"}`)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestTransformPostBody_Malformed(t *testing.T) {
	_, err := transformPostBody([]byte(`{"key": "value", "redirect_uri": "https://example.com/abc?def=ghi"`), "example2.com")
	assert.Error(t, err)
}
