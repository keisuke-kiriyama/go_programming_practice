package main

import (
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	// t, _ := template.ParseFiles("layout.html")
	// t.ExecuteTemplate(w, "layout", "")

	rand.Seed(time.Now().Unix())
	var t *template.Template
	// if rand.Intn(10) > 5 {
	// 	t, _ = template.ParseFiles("layout.html", "red_hello.html")
	// } else {
	// 	t, _ = template.ParseFiles("layout.html", "blue_hello.html")
	// }

	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("layout.html", "red_hello.html")
	} else {
		t, _ = template.ParseFiles("layout.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
