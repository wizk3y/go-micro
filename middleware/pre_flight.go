package middleware

import (
	"net/http"
)

// PreFlightMiddleware is a middleware handler that support pre-flight OPTIONS request
type PreFlightMiddleware struct {
}

// NewPreFlightMiddleware returns a new *Middleware which writes to a given logger.
func NewPreFlightMiddleware() *PreFlightMiddleware {
	return &PreFlightMiddleware{}
}

// ServeHTTP --
func (m *PreFlightMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Method == http.MethodOptions {
		rw.WriteHeader(http.StatusOK)
		return
	}

	next(rw, r)
}
