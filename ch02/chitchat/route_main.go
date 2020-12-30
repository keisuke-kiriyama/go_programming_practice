package main

import (
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.threads()
	if err == nil {
		_, err := session(w, r)
		var templates *template.Template
		if err != nil {
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
