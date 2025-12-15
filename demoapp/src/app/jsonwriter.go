package app

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func writeJson(w http.ResponseWriter, r *http.Request, status int, data interface{}, headers http.Header, log *logrus.Logger) {
	js, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Error while serializing JSON for request %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, "The server could not process your request", http.StatusInternalServerError)
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}