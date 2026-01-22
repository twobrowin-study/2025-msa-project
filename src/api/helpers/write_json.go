package helpers

import (
	"encoding/json"
	"maps"
	"net/http"

	"otus.ru/tbw/msa-25/src/deps"
)

func WriteJson(w http.ResponseWriter, r *http.Request, status int, data any, headers http.Header, deps *deps.Deps) {
	js, err := json.Marshal(data)
	if err != nil {
		deps.Log.Errorf("Error while serializing JSON for request %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, "The server could not process your request", http.StatusInternalServerError)
	}
	maps.Copy(w.Header(), headers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

func WriteJsonError(w http.ResponseWriter, r *http.Request, status int, message string, err error, deps *deps.Deps) {
	deps.Log.Errorf("%s in %s %s request: %v", message, r.Method, r.URL.Path, err)
	WriteJson(w, r, status, map[string]any{"code": status, "message": message}, nil, deps)
}
