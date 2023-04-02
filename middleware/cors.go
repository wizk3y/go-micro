package middleware

import (
	"net/http"
)

// CorsMiddleware is a middleware handler that support CORS response header
type CorsMiddleware struct {
}

// NewCorsMiddleware returns a new *Middleware which writes to a given logger.
func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

// ServeHTTP --
func (m *CorsMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	rw.Header().Add("Access-Control-Allow-Origin", "*")
	rw.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	rw.Header().Add("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Authorization")

	next(rw, r)
}
