package http

import (
	"log"
	"net/http"

	"github.com/julianinsua/codis/internal/database"
)

func (srv *Server) handleUserContent(w http.ResponseWriter, r *http.Request, usr database.User) {
	postList, err := srv.store.GetPostsWithTags(r.Context(), usr.ID)
	if err != nil {
		// exec the error template
		log.Printf("error getting posts: %v", err)
	}

	// ID          uuid.UUID
	// Title       string
	// Description sql.NullString
	// Status      sql.NullString
	// UserID      uuid.UUID
	// PostTags    json.RawMessage
	err = srv.templates["page/userContent"].Execute(w, struct {
		List     []database.PostsView
		Username string
	}{List: postList, Username: usr.Username})
	if err != nil {
		log.Printf("fatal error while answering request on handleContent %s", err)
	}
}
