package app

import (
	"net/http"
	"time"

	"otus.ru/tbw/msa-25/src/deps"
)

func loggingMiddleware(next http.Handler, deps *deps.Deps) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		deps.Log.Infof("%s %s - %s", r.Method, r.URL.Path, time.Since(start))
	})
}
