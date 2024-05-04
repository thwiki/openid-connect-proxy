package main

import (
	"encoding/json"
)

func transformPostBody(body []byte, redirect string) ([]byte, error) {
	var jsonData map[string]interface{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, err
	}

	if redirectURI, ok := jsonData["redirect_uri"].(string); ok {
		jsonData["redirect_uri"] = transformRedirectURI(redirectURI, redirect)
	}

	newBody, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	return newBody, nil
}
