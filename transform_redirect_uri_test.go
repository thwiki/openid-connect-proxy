package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformRedirectURI_Empty(t *testing.T) {
	uri := "https://example.com/abc?def=ghi"
	redirect := ""
	expected := "https://example.com/abc?def=ghi"

	result := transformRedirectURI(uri, redirect)
	assert.Equal(t, expected, result)
}

func TestTransformRedirectURI_Valid(t *testing.T) {
	uri := "https://example.com/abc?def=ghi"
	redirect := "example2.com"
	expected := "https://example2.com/abc?def=ghi"

	result := transformRedirectURI(uri, redirect)
	assert.Equal(t, expected, result)
}

func TestTransformRedirectURI_Invalid(t *testing.T) {
	uri := "invalid-uri"
	redirect := "new-example.com"
	expected := "invalid-uri"

	result := transformRedirectURI(uri, redirect)
	assert.Equal(t, expected, result)
}
