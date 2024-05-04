package main

import (
	"net/url"
)

func transformPostQuery(body []byte, redirect string) ([]byte, error) {
	form, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}

	if form.Has("redirect_uri") {
		form.Set("redirect_uri", transformRedirectURI(form.Get("redirect_uri"), redirect))
	}

	return []byte(form.Encode()), err
}
