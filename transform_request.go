package main

import (
	"bytes"
	"io"
	"net/http"
)

func transformRequest(r *http.Request, redirect string) {
	newQuery := transformGetQuery(r.URL.Query(), redirect)
	r.URL.RawQuery = newQuery

	if r.Method != "POST" {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()

	var newBody []byte
	if r.Header.Get("Content-Type") == "application/json" {
		newBody, err = transformPostBody(body, redirect)
	} else {
		newBody, err = transformPostQuery(body, redirect)
	}

	if err != nil {
		return
	}
	r.Body = io.NopCloser(bytes.NewReader(newBody))
}
