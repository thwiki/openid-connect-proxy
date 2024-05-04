package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformGetQuery_RedirectURI(t *testing.T) {
	query := url.Values{}
	query.Set("redirect_uri", "https://example.com/abc?def=ghi")
	result := transformGetQuery(query, "example2.com")
	expected := "redirect_uri=https%3A%2F%2Fexample2.com%2Fabc%3Fdef%3Dghi"
	assert.Equal(t, expected, result)
}

func TestTransformGetQuery_ScopeProfileEmail(t *testing.T) {
	query := url.Values{}
	query.Set("scope", "profile email")
	result := transformGetQuery(query, "example2.com")
	expected := "scope=basic+mwoauth-authonlyprivate"
	assert.Equal(t, expected, result)
}

func TestTransformGetQuery_ScopeEmailProfile(t *testing.T) {
	query := url.Values{}
	query.Set("scope", "email profile")
	result := transformGetQuery(query, "example2.com")
	expected := "scope=mwoauth-authonlyprivate+basic"
	assert.Equal(t, expected, result)
}

func TestTransformGetQuery_ScopeEmailProfileOther(t *testing.T) {
	query := url.Values{}
	query.Set("scope", "email profile other")
	result := transformGetQuery(query, "example2.com")
	expected := "scope=mwoauth-authonlyprivate+basic"
	assert.Equal(t, expected, result)
}

func TestTransformGetQuery_NoRedirectURI(t *testing.T) {
	query := url.Values{}
	query.Set("other_param", "value")
	result := transformGetQuery(query, "example2.com")
	expected := "other_param=value"
	assert.Equal(t, expected, result)
}
