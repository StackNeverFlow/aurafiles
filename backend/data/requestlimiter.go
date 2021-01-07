package data

import (
	"net/http"
	"time"
)

var (
	limit []string
)

// isRequestLimit is used to check if a ip address has send to many requests
func isRequestLimit(ip string) bool {
	for i := 0; i < len(limit); i++ {
		if limit[i] == ip {
			return true
		}
	}
	return false
}

// setRequestLimit is used to set the request limit and add a ip address to the limit list
func setRequestLimit(ip string) {
	limit = append(limit, ip)
	time.AfterFunc(2*time.Second, func() {
		if isRequestLimit(ip) {
			for i := 0; i < len(limit); i++ {
				if limit[i] == ip {
					limit[i] = ""
				}
			}
		}
	})
}

// CheckRequestLimit is used to check if a ip address has send to many requests
// It also executes a http.Error when someone has already send
func CheckRequestLimit(ip string, w http.ResponseWriter) bool {
	if isRequestLimit(ip) {
		http.Error(w, "to many requests", http.StatusTooManyRequests)
		return false
	} else {
		go setRequestLimit(ip)
		return true
	}
}
