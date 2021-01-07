package data

import (
	"encoding/base64"
	"net/http"
	"strings"
)

const (
	masterUsername = "aura"
	masterPassword = "12345"
)

// Auth is used to authenticate a user using the header and api key
func Auth(w http.ResponseWriter, r *http.Request) bool {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return false
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 || !validate(pair[0], pair[1]) {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return false
	}

	return true
}

// validate is used to validate and check the credentials someone entered
func validate(username, password string) bool {
	return username == masterUsername && password == masterPassword
}
