package http

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"github.com/eknkc/amber"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	db "github.com/julianinsua/codis/database"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/wikilink"
)

const (
	PORT = "8080"
)

type Server struct {
	store     *db.Store
	router    chi.Router
	templates map[string]*template.Template
	mdParser  mdConverter
}

func (srv *Server) Start() {
	srv.compileTemplates()
	srv.setCORSHeaders()
	srv.serveStaticContent()
	srv.setRoutes()

	server := &http.Server{
		Handler: srv.router,
		Addr:    ":" + PORT,
	}

	log.Printf("Server running on port :%v", PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (srv *Server) setCORSHeaders() {
	srv.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*s"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}

func (srv *Server) setRoutes() {
	srv.router.Get("/home", srv.handleHome)
	srv.router.Get("/content", srv.handleArticleList)
	srv.router.Get("/content/{articleName}", srv.handleArticle)
}

func (srv *Server) compileTemplates() {
	temple, err := amber.CompileDir("views", amber.DirOptions{
		Ext:       ".amber",
		Recursive: true,
	}, amber.Options{})

	if err != nil {
		log.Fatal("can't compile templates", err)
	}
	srv.templates = temple
}

func (srv *Server) serveStaticContent() {
	fs := http.FileServer(http.Dir("static"))
	srv.router.Handle("/static/*", http.StripPrefix("/static/", fs))
}

func NewServer(store *db.Store) *Server {
	router := chi.NewRouter()
	mdParser := NewMdParser()
	return &Server{router: router, store: store, mdParser: mdParser}
}

type mdConverter interface {
	Convert([]byte) (template.HTML, error)
}

type MdParser struct {
}

func (mdp MdParser) Convert(mdFile []byte) (template.HTML, error) {
	var buffer bytes.Buffer
	err := goldmark.New(
		goldmark.WithExtensions(extension.GFM, &wikilink.Extender{}),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	).Convert(mdFile, &buffer)
	if err != nil {
		return template.HTML(""), err
	}
	return template.HTML(buffer.String()), nil
}

func NewMdParser() mdConverter {
	return MdParser{}
}
