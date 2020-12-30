package main

import (
	"fmt"
	"net/http"
)

// これはハンドラ関数→ハンドラのように振舞う関数
// ServeHTTPと同じシグネチャをもつ
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// 適切なシグネチャを持った関数fから、メソッドfをもつハンドラに変換してくれる
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
