package app

import (
	"net/http"

	"otus.ru/tbw/msa-25/src/api"
	"otus.ru/tbw/msa-25/src/api/user"
	"otus.ru/tbw/msa-25/src/deps"
)

func newRouter(deps *deps.Deps) *http.ServeMux {
	router := http.NewServeMux()

	api.MuxRegisterPath(router, "GET", "/health", "", func(w http.ResponseWriter, r *http.Request) {
		api.WriteJson(w, r, http.StatusOK, map[string]string{"status": "OK"}, nil, deps)
	})

	router.Handle("/user/", user.RegisterRoutes("/user", deps))

	return router
}
