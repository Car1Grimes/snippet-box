package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Car1Grimes/snippet-box/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	snippets, err := app.snippets.Recent()
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.serverError(w, r, models.ErrNoRecord)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	app.render(w, r, http.StatusOK, "home.tmpl.html", templateData{
		Snippets: snippets,
	})
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	app.render(w, r, http.StatusOK, "view.tmpl.html", templateData{
		Snippet: snippet,
	})
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

func (app *application) snippetViewRecent(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Recent()
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.serverError(w, r, models.ErrNoRecord)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	for _, val := range snippets {
		fmt.Fprintf(w, "%v", val)
	}
}
