package main

import (
	"net/http"
	"text/template"
)

// 条件アクション
// func process(w http.ResponseWriter, r *http.Request) {
// 	t, _ := template.ParseFiles("tmpl.html")
// 	rand.Seed(time.Now().Unix())
// 	t.Execute(w, rand.Intn(10) > 5)
// }

// イテレータアクション
// func process(w http.ResponseWriter, r *http.Request) {
// 	t, _ := template.ParseFiles("tmpl.html")
// 	daysOfWeek := []string{"月", "火", "水", "木", "金", "土", "日"}
// 	t.Execute(w, daysOfWeek)
// }

// 代入アクション
func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "hello")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
