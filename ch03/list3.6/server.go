package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

// ServeHTTPを持ったインターフェース→ハンドラ
// このメソッドは、http.ResponseWriterと構造体Requestへのポインタという２つの引数をとる
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
