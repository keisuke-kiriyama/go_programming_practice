package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// ずべての場所から参照可能でなければいけないため、大文字
type Post struct {
	XMLName  xml.Name  `xml:"post"`     // XML要素名自体を保存する
	Id       string    `xml:"id, attr"` // attrがついているのは、XMLの属性
	Content  string    `xml:"content"`  // フィールド名と同じ名前のXML要素と紐付けられる
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"` // XML要素から未処理のままで生のXMLを得る
	Comments []Comment `xml:"comments>comment"`
}

// データを表す構造体を定義する
type Author struct {
	Id   string `xml:"id,attr"`   // 属性Idを取り込むため、別構造体で定義
	Name string `xml:",chardata"` // XML要素の文字データの保存
}

type Comment struct {
	Id      string `xml:"id, attr"`
	Content string `xml:"content"`
	Author  Author `xml: "author"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Erro opening XML file:", err)
		return
	}

	// 小さめのXMLの場合
	// defer xmlFile.Close()
	// xmlData, err := ioutil.ReadAll(xmlFile)
	// if err != nil {
	// 	fmt.Println("Error reading XML data:", err)
	// }

	// var post Post
	// xml.Unmarshal(xmlData, &post)
	// fmt.Println(post)

	// 大きいXMLやストリームとして入力されるXMLの場合
	// XMLからdecoderを生成
	decoder := xml.NewDecoder(xmlFile)
	// decoder内のXMLデータを順次処理
	for {
		// 各処理でdecoderからトークンを取得
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}

		// トークンの型をチェック
		switch se := t.(type) {
		case xml.StartElement: // XML要素の開始タグ
			if se.Name.Local == "comment" {
				var comment Comment
				// XMLデータを構造デコード
				decoder.DecodeElement(&comment, &se)
			}
		}
	}
}
