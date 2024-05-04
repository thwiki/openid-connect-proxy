package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port       = flag.Int("port", 8000, "int listening port")
	redirect   = flag.String("redirect", "", "string redirect host")
	upstream   = flag.String("upstream", "", "string upstream host")
	proxyPaths = map[string]string{
		"/protocol/openid-connect/auth":     "/rest/oauth2/authorize",
		"/protocol/openid-connect/token":    "/rest/oauth2/access_token",
		"/protocol/openid-connect/userinfo": "/rest/oauth2/resource/profile",
	}
	responseKeyMap = map[string]string{
		"username":        "name",
		"realname":        "preferred_username",
		"confirmed_email": "email_verified",
	}
)

func main() {
	flag.Parse()

	r := mux.NewRouter()
	for k, v := range proxyPaths {
		r.HandleFunc(k, func(w http.ResponseWriter, r *http.Request) {
			transformRequest(r, *redirect)
			body, err := proxyRequest(v, w, r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			newBody, err := transformResponse(body)
			if err == nil {
				w.Write(newBody)
			} else {
				w.Write(body)
			}
		})
	}

	fmt.Printf("Listening on port %d...\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), r)
}

func proxyRequest(upstreamPath string, w http.ResponseWriter, r *http.Request) ([]byte, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	upstreamURL := r.URL
	upstreamURL.Scheme = "https"
	upstreamURL.Host = *upstream
	upstreamURL.Path = upstreamPath
	req, err := http.NewRequest(r.Method, upstreamURL.String(), r.Body)
	if err != nil {
		return nil, err
	}

	req.Header = r.Header.Clone()
	req.Header.Del("Accept-Encoding")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
