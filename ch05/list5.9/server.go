package main

import (
	"html/template"
	"net/http"
)

// インクルードアクション
func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html") // 解析するテンプレート全て与える。1番目のテンプレートがメインテンプレート
	t.Execute(w, "Hello World!")                      // Executeでメインテンプレートが呼び出される
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
