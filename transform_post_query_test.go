package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformPostQuery_Empty(t *testing.T) {
	body := []byte("")
	redirect := ""
	expected := []byte("")

	result, err := transformPostQuery(body, redirect)
	assert.NoError(t, err)
	assert.Equal(t, (expected), (result))
}

func TestTransformPostQuery_EmptyRedirect(t *testing.T) {
	body := []byte("redirect_uri=https%3A%2F%2Fexample.com%2Fabc%3Fdef%3Dghi")
	redirect := ""
	expected := []byte("redirect_uri=https%3A%2F%2Fexample.com%2Fabc%3Fdef%3Dghi")

	result, err := transformPostQuery(body, redirect)
	assert.NoError(t, err)
	assert.Equal(t, (expected), (result))
}

func TestTransformPostQuery_Valid(t *testing.T) {
	body := []byte("redirect_uri=https%3A%2F%2Fexample.com%2Fabc%3Fdef%3Dghi")
	redirect := "example2.com"
	expected := []byte("redirect_uri=https%3A%2F%2Fexample2.com%2Fabc%3Fdef%3Dghi")

	result, err := transformPostQuery(body, redirect)
	assert.NoError(t, err)
	assert.Equal(t, (expected), (result))
}

func TestTransformPostQuery_Invalid(t *testing.T) {
	body := []byte("redirect_uri=invalid-uri")
	redirect := "example2.com"
	expected := []byte("redirect_uri=invalid-uri")

	result, err := transformPostQuery(body, redirect)
	assert.NoError(t, err)
	assert.Equal(t, (expected), (result))
}
