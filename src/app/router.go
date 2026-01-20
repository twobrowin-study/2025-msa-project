package app

import (
	"net/http"

	"otus.ru/tbw/msa-25/src/api/helpers"
	"otus.ru/tbw/msa-25/src/api/user"
	"otus.ru/tbw/msa-25/src/deps"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func newRouter(deps *deps.Deps) *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /metrics/", promhttp.HandlerFor(deps.Prometheus.Registry, promhttp.HandlerOpts{}))

	helpers.MuxRegisterPath(router, "GET", "/health", "", func(w http.ResponseWriter, r *http.Request) {
		helpers.WriteJson(w, r, http.StatusOK, map[string]string{"status": "OK"}, nil, deps)
	})

	router.Handle("/user/", deps.Prometheus.WrapHandler("/user", user.RegisterRoutes("/user", deps)))

	return router
}
