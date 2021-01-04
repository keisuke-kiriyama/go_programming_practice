// testパッケージの使用
// package main

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"strings"
// 	"testing"
// )

// var mux *http.ServeMux
// var writer *httptest.ResponseRecorder

// func TestMain(m *testing.M) {
// 	setUp()
// 	code := m.Run()
// 	os.Exit(code)
// }

// func setUp() {
// 	// テストを実行するマルチプレクサを生成
// 	mux = http.NewServeMux()
// 	// テスト対象のハンドラを付加
// 	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
// 	// 返されたHTTPレスポンスを取得
// 	writer = httptest.NewRecorder()
// }

// func TestHandleGet(t *testing.T) {
// 	// テストしたいハンドラ宛てのリクエストを作成
// 	request, _ := http.NewRequest("GET", "/post/1", nil)
// 	// テスト対象のハンドラにリクエストを送信
// 	mux.ServeHTTP(writer, request)

// 	if writer.Code != 200 {
// 		t.Errorf("Response code is %v", writer.Code)
// 	}
// 	var post Post
// 	json.Unmarshal(writer.Body.Bytes(), &post)
// 	if post.Id != 1 {
// 		t.Error("Cannot retrieve JSON post")
// 	}
// }

// func TestHandlePut(t *testing.T) {
// 	json := strings.NewReader(`{"content":"updated post", "author": "Sau Sheong`)
// 	request, _ := http.NewRequest("PUT", "/post/1", json)
// 	mux.ServeHTTP(writer, request)
// 	if writer.Code != 200 {
// 		t.Errorf("Response code is %v", writer.Code)
// 	}
// }

// checkの使用
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	// エクスポートされた識別子はパッケージ名を省略してアクセスできる
	. "gopkg.in/check.v1"
)

// テストスイートの作成
type PostTestSuite struct {
	mux    *http.ServeMux
	post   *FakePost
	writer *httptest.ResponseRecorder
}

func init() {
	// テストスイートの登録
	Suite(&PostTestSuite{})
}

// パッケージtestingとの結合
func Test(t *testing.T) { TestingT(t) }

func (s *PostTestSuite) SetUpTest(c *C) {
	s.post = &FakePost{}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/post/", handleRequest(s.post))
	s.writer = httptest.NewRecorder()
}

func (s *PostTestSuite) TestGetPost(c *C) {
	request, _ := http.NewRequest("GET", "/post/1", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	var post Post
	json.Unmarshal(s.writer.Body.Bytes(), &post)
	c.Check(post.Id, Equals, 1)
}

func (s *PostTestSuite) TestPutPost(c *C) {
	json := strings.NewReader(`{"content":"updated post", "author": "Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	c.Check(s.post.Id, Equals, 1)
	c.Check(s.post.Content, Equals, "updated post")
}
