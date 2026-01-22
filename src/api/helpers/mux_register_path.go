package helpers

import (
	"fmt"
	"net/http"
)

func MuxRegisterPath(mux *http.ServeMux, method string, pathPrefix string, pathSuffix string, handler func(http.ResponseWriter, *http.Request)) {
	pattern := fmt.Sprintf("%s %s%s", method, pathPrefix, pathSuffix)
	mux.HandleFunc(pattern, handler)
	mux.HandleFunc(pattern+"/", handler)
}
