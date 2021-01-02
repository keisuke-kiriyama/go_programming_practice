package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	// テストを実行するマルチプレクサを生成
	mux = http.NewServeMux()
	// テスト対象のハンドラを付加
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
	// 返されたHTTPレスポンスを取得
	writer = httptest.NewRecorder()
}

func TestHandleGet(t *testing.T) {
	// テストしたいハンドラ宛てのリクエストを作成
	request, _ := http.NewRequest("GET", "/post/1", nil)
	// テスト対象のハンドラにリクエストを送信
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	json := strings.NewReader(`{"content":"updated post", "author": "Sau Sheong`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
