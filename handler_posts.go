package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/imhasandl/go-restapi/internal/database"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	Body      string    `json:"body"`
}

func (cfg *apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserID uuid.UUID `json:"user_id"`
		Body   string    `json:"body"`
	}
	type responce struct {
		Post
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "can't decode the body", err)
		return
	}

	post, err := cfg.db.CreatePost(r.Context(), database.CreatePostParams{
		ID:     uuid.New(),
		UserID: params.UserID,
		Body:   params.Body,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "can't post it", err)
		return
	}

	respondWithJSON(w, http.StatusOK, responce{
		Post: Post{
			ID: post.ID,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			UserID: post.UserID,
			Body: post.Body,
		},
	})
}

func (cfg *apiConfig) handlerListPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := cfg.db.GetPosts(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "can't list the posts", err)
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
