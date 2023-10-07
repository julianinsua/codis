package api

import (
	"html/template"
	"log"
	"net/http"

	"github.com/eknkc/amber"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	db "github.com/julianinsua/codis/database"
)

const (
	PORT = "8080"
)

type Server struct {
	store     *db.Store
	router    chi.Router
	templates map[string]*template.Template
}

func (srv *Server) Start() {
	srv.compileTemplates()
	srv.setCORSHeaders()
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
	fs := http.FileServer(http.Dir("static"))
	srv.router.Handle("/static/*", http.StripPrefix("/static/", fs))
	srv.router.Get("/home", srv.handleHome)
}

func (srv *Server) compileTemplates() {
	temple, err := amber.CompileDir("views", amber.DirOptions{
		Ext:       ".amber",
		Recursive: true,
	}, amber.Options{})

	if err != nil {
		log.Fatal("can't compile templates")
	}
	srv.templates = temple
}

func (srv *Server) serveStaticContent() {
}

func NewServer(store *db.Store) *Server {
	router := chi.NewRouter()
	return &Server{router: router, store: store}
}
