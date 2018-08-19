package ravenchi_test

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/loikg/ravenchi"
)

func ExampleSentryRecovery() {
	r := chi.NewRouter()

	// Apply the middleware to the router
	r.Use(ravenchi.SentryRecovery)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		panic("catched")
	})
}
