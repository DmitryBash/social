package main

import (
	"net/http"
)

func (app *application) GetUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	feed, err := app.store.Posts.GetUserFeed(ctx, int64(3))

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
	}
}
