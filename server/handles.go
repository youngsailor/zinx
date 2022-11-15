package server

import (
	"net/http"
)

func WsHandler(w http.ResponseWriter, r *http.Request) {
	GWServer.Serve(w, r)
}
