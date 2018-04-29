package web

import (
	"log"
	"net/http"
)

// Adapter functions type
type Adapter func(h http.Handler) http.Handler

// Adapt apply all adapters
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

// Notify prints timestamps before and after controller call
func Notify(logger *log.Logger) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println("before")
			defer logger.Println("after")
			h.ServeHTTP(w, r)
		})
	}
}
