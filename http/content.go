package http

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (srv *Server) handleArticleList(w http.ResponseWriter, r *http.Request) {
	data := struct{ Content template.HTML }{Content: template.HTML("<b>pepito</b>")}
	err := srv.templates["page/content"].Execute(w, data)
	if err != nil {
		log.Printf("fatal error while answering request on handleArticleList: %s", err)
	}
}

func (srv *Server) handleArticle(w http.ResponseWriter, r *http.Request) {
	articleName := chi.URLParam(r, "articleName")
	matched := strings.Split(articleName, ".")[0]
	mdArticle, err := srv.files.GetFile(matched)
	if err != nil {
		log.Printf("unable to read markdown file: %s", err)
		data := struct {
			Code   int
			Legend string
		}{404, "The resource you want isn't where it's supposed to be"}
		err = srv.templates["page/error"].Execute(w, data)
		if err != nil {
			log.Printf("unable to respond to article request: %s", err)
		}
		return
	}

	// Parse heading ids, parser auto id's in kebab-case all headings

	htmlTempl, err := srv.mdParser.Convert(mdArticle)
	if err != nil {
		log.Printf("Error parsing MD: %v", err)
	}

	data := struct{ Content template.HTML }{Content: htmlTempl}
	err = srv.templates["page/content"].Execute(w, data)
	if err != nil {
		log.Printf("fatal error while answering request on handleArticle %s", err)
	}
}
