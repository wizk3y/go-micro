package middleware

import (
	"net/http"
	"runtime"

	"github.com/urfave/negroni"
	"github.com/wizk3y/go-micro/logger"
)

// RecoveryMiddleware is a middleware that recovers from any panics and writes a 500 if there was one.
type RecoveryMiddleware struct {
	PrintStack bool
	StackAll   bool
	StackSize  int
	Formatter  negroni.PanicFormatter
}

func NewRecoveryMiddleware(printStack bool) *RecoveryMiddleware {
	return &RecoveryMiddleware{
		PrintStack: printStack,
		StackAll:   false,
		StackSize:  1024 * 8,
		Formatter:  &negroni.TextPanicFormatter{},
	}
}

// ServeHTTP --
func (m *RecoveryMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)

			stack := make([]byte, m.StackSize)
			if p, ok := err.(*timeoutInnerPanic); ok && p != nil {
				stack = p.Stack
			} else {
				stack = stack[:runtime.Stack(stack, m.StackAll)]
			}
			infos := &negroni.PanicInformation{RecoveredPanic: err, Request: r}

			if m.PrintStack {
				infos.Stack = stack
			}
			logger.Infow("Recovery catch panic", "error", err, "stacktrace", string(stack))
			m.Formatter.FormatPanicError(rw, r, infos)
		}
	}()

	next(rw, r)
}
