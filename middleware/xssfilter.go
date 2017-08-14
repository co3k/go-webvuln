package middleware

import (
	"net/http"
)

func DisableXSSFilter(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "0")

		f(w, r)
	}
}
