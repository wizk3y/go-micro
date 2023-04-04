package middleware

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

// TimeoutMiddleware is a middleware handler that support CORS response header
type TimeoutMiddleware struct {
	timeout   time.Duration
	stackAll  bool
	stackSize int
}

// NewTimeoutMiddleware returns a new *Middleware which writes to a given logger.
func NewTimeoutMiddleware(timeout time.Duration) *TimeoutMiddleware {
	return &TimeoutMiddleware{
		timeout:   timeout,
		stackAll:  false,
		stackSize: 1024 * 8,
	}
}

// ServeHTTP --
func (m *TimeoutMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), m.timeout)
	defer cancelCtx()

	r = r.WithContext(ctx)

	done := make(chan struct{})
	tw := &timeoutWriter{}                 // prevent double write to response
	panicChan := make(chan interface{}, 1) // catch panic in go routine and throw to main process
	go func() {
		defer func() {
			if err := recover(); err != nil {
				stack := make([]byte, m.stackSize)
				stack = stack[:runtime.Stack(stack, m.stackAll)]

				panicChan <- &timeoutInnerPanic{err: err, Stack: stack}
			}
		}()

		next(tw, r)
		close(done)
	}()
	select {
	case p := <-panicChan:
		panic(p)
	case <-done:
		tw.mu.Lock()
		defer tw.mu.Unlock()
		dst := rw.Header()
		for k, vv := range tw.h {
			dst[k] = vv
		}
		if !tw.wroteHeader {
			tw.code = http.StatusOK
		}
		rw.WriteHeader(tw.code)
		rw.Write(tw.wbuf.Bytes())
	case <-ctx.Done():
		rw.WriteHeader(http.StatusServiceUnavailable)
		rw.Write([]byte("Request Timed-out"))
		tw.timedOut = true
	}
}

type timeoutWriter struct {
	mu          sync.Mutex
	code        int
	h           http.Header
	wbuf        bytes.Buffer
	timedOut    bool
	wroteHeader bool
}

func (tw *timeoutWriter) Header() http.Header {
	if tw.h == nil {
		tw.h = make(http.Header)
	}

	return tw.h
}

func (tw *timeoutWriter) Write(p []byte) (int, error) {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	if tw.timedOut {
		return 0, context.DeadlineExceeded
	}
	if !tw.wroteHeader {
		tw.writeHeader(http.StatusOK)
	}
	return tw.wbuf.Write(p)
}

func (tw *timeoutWriter) WriteHeader(code int) {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	if tw.timedOut || tw.wroteHeader {
		return
	}
	tw.writeHeader(code)
}

func (tw *timeoutWriter) writeHeader(code int) {
	tw.wroteHeader = true
	tw.code = code
}

type timeoutInnerPanic struct {
	err   interface{}
	Stack []byte
}

func (p *timeoutInnerPanic) Error() string {
	return fmt.Sprintf("%s", p.err)
}
