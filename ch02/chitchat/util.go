package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func session(w http.ResponseWriter, r *http.Request) (sess data.Sesstion, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Sesstion{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.new("Invalid session")
		}
	}
	return
}

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", files))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
