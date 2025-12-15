package app

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func newRouter(log *logrus.Logger) *http.ServeMux {
    router := http.NewServeMux()

    router.HandleFunc("GET /health/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Переделать на middleware чтобы логгировать все запросы универсально
		log.Infof("%s %s", r.Method, r.URL.Path)
		writeJson(w, r, http.StatusOK, map[string]string{"status": "OK"}, nil, log)
    })

    return router
}