package api

import (
	"net/http"
)

func (srv *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	// util.RespondWithJson(w, 200, "Hello world!")
	data := struct{ Key string }{
		Key: "value",
	}
	srv.templates["page/home"].Execute(w, data)
}
