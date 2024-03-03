package http

import (
	"log"
	"net/http"
)

func (srv *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	// util.RespondWithJson(w, 200, "Hello world!")
	data := struct{ Key string }{
		Key: "value",
	}
	err := srv.templates["page/home"].Execute(w, data)
	if err != nil {
		log.Fatal("fatal error while answering request on handleHome")
	}
}
