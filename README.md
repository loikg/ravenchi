# ravenchi
Ravenchi implement a middleware for report panic to sentry.io using the raven-go client. It also integrate with chi middleware ecosystem by logging any appropriate informations.

`raven-go` client must be [initialize with `SetDSN()` or by using environment variables](https://docs.sentry.io/clients/go/integrations/http/)

Here is an example on how to use the middleware.
```
r := chi.NewRouter()

// Apply the middleware to the router
r.Use(ravenchi.SentryRecovery)

r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    panic("catched")
})
```
