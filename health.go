package util

import "net/http"

func init() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		return
	})
}
