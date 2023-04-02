package middleware

import (
	"github.com/urfave/negroni"
	"github.com/wizk3y/go-micro/logger"
)

// NewRecoveryMiddleware --
func NewRecoveryMiddleware(printStack bool) *negroni.Recovery {
	recovery := negroni.NewRecovery()
	recovery.ErrorHandlerFunc = recoveryErrorHandlerFunc
	recovery.PrintStack = printStack
	return recovery
}

func recoveryErrorHandlerFunc(error interface{}) {
	logger.Errorw("Recovery catch panic", "error", error)
}
