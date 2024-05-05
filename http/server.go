package http

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	katex "github.com/FurqanSoftware/goldmark-katex"
	callout "github.com/VojtaStruhar/goldmark-obsidian-callout"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/eknkc/amber"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	db "github.com/julianinsua/codis/database"
	"github.com/julianinsua/codis/util"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/mermaid"
	"go.abhg.dev/goldmark/wikilink"
)

const (
	PORT = "8080"
)

type Server struct {
	store     *db.Store
	router    chi.Router
	templates map[string]*template.Template
	mdParser  MdConverter
	files     util.FileManager
	config    util.Config
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
		AllowedOrigins:   []string{"https://*", "http://*"},
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
	srv.router.Post("/content", srv.handleCreateArticle)
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

func NewServer(store *db.Store, mdParser MdConverter, config util.Config) *Server {
	router := chi.NewRouter()

	return &Server{router: router, store: store, mdParser: mdParser, config: config}
}

type MdConverter interface {
	Convert([]byte) (template.HTML, error)
}

type MdParser struct {
}

func (mdp MdParser) Convert(mdFile []byte) (template.HTML, error) {
	var buffer bytes.Buffer
	err := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			&wikilink.Extender{},
			highlighting.NewHighlighting(
				highlighting.WithStyle("catppuccin-mocha"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
			&mermaid.Extender{},
			&katex.Extender{},
			callout.ObsidianCallout,
		),
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

func NewMdParser() MdConverter {
	return MdParser{}
}
