package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type post struct {
	ID string
	title string
    description string
    likes string
    comments string
    shares string
    created_at time.Time
}

type postHandler struct {
	db *sql.DB
	//more dps here...
}

//contructor fn
func NewPostHandler(db *sql.DB) *postHandler {
	return &postHandler{
		db: db,
	}
}

func (p postHandler) Post(w http.ResponseWriter, r *http.Request) {
	rows, err := p.db.Query(
		`SELECT id, title, description, likes, shares, comments, created_at
		FROM post
		ORDER BY created_at DESC
		LIMIT 100`,
		)

	if err != nil {
		log.Printf("query %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}

	defer rows.Close()

	posts := []post{}
	
	for rows.Next() {
		var p post
		if err := rows.Scan(&p.ID, &p.title, &p.description, &p.likes, &p.shares, &p.comments, &p.created_at); err != nil {
			log.Printf("rows.Scan: %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows.err: %v:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	_ = json.NewEncoder(w).Encode(posts)	
}

func (p postHandler) Delete (w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_ = ctx.Value("requestCtxId").(string)
	id := r.PathValue("id")
	
	
	_, err := p.db.Exec(
		`DELETE FROM post WHERE id = $1`, id,
	)

	if err != nil {
		log.Printf("delete: %v", err)
		//can log requestId here using slog
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
} 