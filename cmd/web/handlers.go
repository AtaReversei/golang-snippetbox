package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errLog.Println(err.Error())
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.errLog.Println(err.Error())
		app.serverError(w, err)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "showing snippet with the id of %d", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		//w.WriteHeader(405)
		//w.Write([]byte("Method not allowed..."))

		//http.Error(w, "Method not allowed", 405)

		app.clientError(w, http.StatusMethodNotAllowed)

		return
	}
	w.Write([]byte("creating snippets"))
}
