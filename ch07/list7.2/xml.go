package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// ずべての場所から参照可能でなければいけないため、大文字
type Post struct {
	XMLName xml.Name `xml:"post"`     // XML要素名自体を保存する
	Id      string   `xml:"id, attr"` // attrがついているのは、XMLの属性
	Content string   `xml:"content"`  // フィールド名と同じ名前のXML要素と紐付けられる
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"` // XML要素から未処理のままで生のXMLを得る
}

// データを表す構造体を定義する
type Author struct {
	Id   string `xml:"id,attr"`   // 属性Idを取り込むため、別構造体で定義
	Name string `xml:",chardata"` // XML要素の文字データの保存
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Erro opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
	}

	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}
