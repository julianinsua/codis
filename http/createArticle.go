package http

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/julianinsua/codis/internal/database"
)

func (srv *Server) handleCreateArticle(w http.ResponseWriter, r *http.Request, usr database.User) {
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
	filename := fmt.Sprintf("%v - %v", usr.Username, fileHeader.Filename)
	storagePath, err := srv.files.SaveFile(file, filename)
	if err != nil {
		log.Printf("error saving file: %v", err)
	}

	createPostParams := database.CreatePostTxParams{
		Title:       "Title",
		Description: sql.NullString{String: "Description", Valid: true},
		Status:      sql.NullString{String: "Published", Valid: true},
		UserID:      usr.ID,
		Path:        storagePath,
		TagNames:    []string{"test_tag_1"},
	}
	// success: Create database entry
	post, err := srv.store.CreatePostTx(r.Context(), createPostParams)
	if err != nil {
		log.Print(post)
		log.Printf("sqlerror: %v", err)
	}
}
