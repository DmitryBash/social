package main

import (
	"net/http"
	"social/internal/store"
)

type CreatePostPayload struct {
	Title   string   `json:title`
	Content string   `json:content`
	Tags    []string `json:tags`
}

func (app *application) createPostsHandler(w http.ResponseWriter, r *http.Request) {
	var paylod CreatePostPayload
	if err := readJSON(w, r, &paylod); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	post := &store.Post{
		Title:   paylod.Title,
		Content: paylod.Content,
		Tags:    paylod.Tags,
		UserID:  1,
	}
	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
