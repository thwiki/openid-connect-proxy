package main

import (
	"net/url"
)

func transformRedirectURI(uri string, redirect string) string {
	if redirect == "" {
		return uri
	}

	u, err := url.ParseRequestURI(uri)

	if err != nil {
		return uri
	}

	u.Host = redirect
	return u.String()
}
