package main

import (
	"encoding/json"
)

func transformResponse(body []byte) ([]byte, error) {
	var jsonData map[string]interface{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, err
	}

	if blocked, ok := jsonData["blocked"].(bool); ok && blocked {
		jsonData = make(map[string]interface{})
	} else {
		for oldKey, newKey := range responseKeyMap {
			if value, ok := jsonData[oldKey]; ok {
				jsonData[newKey] = value
			}
		}
	}

	newBody, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	return newBody, nil
}
