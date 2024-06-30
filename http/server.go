package http

import (
	"html/template"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/eknkc/amber"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/julianinsua/codis/internal/database"
	"github.com/julianinsua/codis/token"
	"github.com/julianinsua/codis/util"
)

// FileManager allows saving and reading files.
type FileManager interface {
	SaveFile(file multipart.File, filename string) (string, error)
	GetFile(filename string) ([]byte, error)
}

// Converter turns a read Markdown file into an HTMl template.
type Converter interface {
	Convert([]byte) (template.HTML, error)
}

type Server struct {
	store      database.Store
	router     chi.Router
	templates  map[string]*template.Template
	mdParser   Converter
	files      FileManager
	config     util.Config
	tokenMaker token.Maker
}

// Starts initiates the servers and listens to the preconfigured port.
func (srv *Server) Start() {
	srv.compileTemplates()
	srv.setCORSHeaders()
	srv.serveStaticContent()
	srv.setRoutes()
	srv.setAuthorizedRoutes()

	server := &http.Server{
		Handler: srv.router,
		Addr:    srv.config.PORT,
	}

	log.Printf("Server running on port :%v", srv.config.PORT)
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
	srv.router.Post("/login", srv.handleLogin)
	srv.router.Post("/signup", srv.createUser)
}

func (srv *Server) setAuthorizedRoutes() {
	srv.router.Post("/content", srv.authorizedHandler(srv.handleCreateArticle))
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

// NewServer creates a new server instance
func NewServer(store database.Store, mdParser Converter, config util.Config, tokenMaker token.Maker, files FileManager) *Server {
	router := chi.NewRouter()

	return &Server{
		router:     router,
		store:      store,
		mdParser:   mdParser,
		config:     config,
		tokenMaker: tokenMaker,
		files:      files,
	}
}
