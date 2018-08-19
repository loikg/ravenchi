// Package ravenchi implement a middleware for report panic to sentry.io.
// It also integrate with chi middleware ecosystem by logging any appropriate
// informations
package ravenchi

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	raven "github.com/getsentry/raven-go"
	"github.com/go-chi/chi/middleware"
)

// SentryRecovery recover from panic, report the error back to sentry, log it,
// and return an internal server error back to the user.
func SentryRecovery(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rval := recover(); rval != nil {

				logEntry := middleware.GetLogEntry(r)
				if logEntry != nil {
					logEntry.Panic(rval, debug.Stack())
				} else {
					fmt.Fprintf(os.Stderr, "Panic: %+v\n", rval)
					debug.PrintStack()
				}

				rvalStr := fmt.Sprint(rval)
				packet := raven.NewPacket(rvalStr, raven.NewException(errors.New(rvalStr), raven.GetOrNewStacktrace(rval.(error), 2, 3, nil)), raven.NewHttp(r))
				raven.Capture(packet, nil)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		handler.ServeHTTP(w, r)
	})
}
