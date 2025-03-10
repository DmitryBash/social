package main

import (
	"net/http"
	"social/internal/store"
	"strconv"

	"errors"

	"github.com/go-chi/chi/v5"
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

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ctx := r.Context()
	post, err := app.store.Posts.GetByID(ctx, id)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			writeJSONError(w, http.StatusNotFound, err.Error())
		default:
			writeJSONError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
