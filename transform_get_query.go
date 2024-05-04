package main

import (
	"net/url"
	"strings"
)

func transformGetQuery(query url.Values, redirect string) string {
	if query.Has("redirect_uri") {
		query.Set("redirect_uri", transformRedirectURI(query.Get("redirect_uri"), redirect))
	}

	if scopes, ok := query["scope"]; ok {
		newScopes := []string{}
		for _, scope := range strings.Split(scopes[0], " ") {
			switch scope {
			case "profile":
				newScopes = append(newScopes, "basic")
			case "email":
				newScopes = append(newScopes, "mwoauth-authonlyprivate")
			}
		}
		query.Set("scope", strings.Join(newScopes, " "))
	}

	return query.Encode()
}
