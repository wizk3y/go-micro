package middleware

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/urfave/negroni"
	"github.com/wizk3y/go-micro/logger"
	"github.com/wizk3y/go-micro/util"
)

// BeforeFunc is the func type used to modify or replace the *logrus.Entry prior
// to calling the next func in the middleware chain
type BeforeFunc func(*http.Request, string, string)

// AfterFunc is the func type used to modify or replace the *logrus.Entry after
// calling the next func in the middleware chain
type AfterFunc func(negroni.ResponseWriter, time.Duration, string)

// DefaultBefore is the default func assigned to *Middleware.Before
func DefaultBefore(req *http.Request, remoteAddr, msg string) {
	logger.Infow(msg,
		"remote", remoteAddr,
	)
}

// DefaultAfter -- is the default func assigned to *Middleware.After
func DefaultAfter(res negroni.ResponseWriter, latency time.Duration, msg string) {
	logger.Infow(msg,
		"status", res.Status(),
		"text_status", http.StatusText(res.Status()),
		"took", latency,
	)
}

// HttpLogMiddleware is a middleware handler that logs the request as it goes in and the response as it goes out.
type HttpLogMiddleware struct {
	Before      BeforeFunc
	After       AfterFunc
	logStarting bool

	// Exclude URLs from logging
	excludeURLs []string
}

// NewHttpLogMiddleware returns a new *Middleware which writes to a given logger.
func NewHttpLogMiddleware() *HttpLogMiddleware {
	return &HttpLogMiddleware{
		Before: DefaultBefore,
		After:  DefaultAfter,
	}
}

// ExcludeURL -- adds a new URL u to be ignored during logging. The URL u is parsed, hence the returned error
func (m *HttpLogMiddleware) ExcludeURL(u string) error {
	if _, err := url.Parse(u); err != nil {
		return err
	}
	m.excludeURLs = append(m.excludeURLs, u)
	return nil
}

// ExcludedURLs -- returns the list of excluded URLs for this middleware
func (m *HttpLogMiddleware) ExcludedURLs() []string {
	return m.excludeURLs
}

// ServeHTTP --
func (m *HttpLogMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if m.Before == nil {
		m.Before = DefaultBefore
	}

	if m.After == nil {
		m.After = DefaultAfter
	}

	for _, u := range m.excludeURLs {
		if r.URL.Path == u {
			next(rw, r)
			return
		}
	}

	start := time.Now()

	// Try to get the real IP
	remoteAddr := util.GetClientAddr(r)

	m.Before(r, remoteAddr, fmt.Sprintf("Started handling api request %v", r.RequestURI))

	next(rw, r)
	latency := time.Since(start)
	res := rw.(negroni.ResponseWriter)

	m.After(res, latency, fmt.Sprintf("Completed handling api request %v", r.RequestURI))
}
