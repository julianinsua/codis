package http

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/julianinsua/codis/internal/database"
)

func (srv *Server) handleCreateArticle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 * 2^20
	if err != nil {
		log.Printf("unable to parse multipart form")
		return
	}

	file, fileHeader, err := r.FormFile("myFile") // this calls ParseMultipartForm automatically
	if err != nil {
		log.Printf("unable to get form file: %s", err)
		return
	}
	defer file.Close()

	log.Printf("file received: %v", fileHeader.Filename)

	// TODO: create a service on the server to store uploaded file
	localFilename, err := srv.files.SaveFile(file, fileHeader.Filename)
	if err != nil {
		log.Printf("error saving file: %v", err)
	}

	// TODO: filename is missing on DB
	createPostParams := database.CreatePostParams{
		Title:       "pepito",
		Description: sql.NullString{String: "null this", Valid: true},
		Status:      sql.NullString{String: "published", Valid: true},
		UserID:      uuid.New(),
		Path:        localFilename,
	}
	// success: Create database entry
	post, err := srv.store.Queries.CreatePost(r.Context(), createPostParams)
	if err != nil {
		log.Print(post)
		log.Printf("sqlerror: %v", err)
	}
}
