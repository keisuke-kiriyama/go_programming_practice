package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id, attr"`
	Content string   `xml:"content"`
	Author  Author   `xml: "author`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml: ",chardata"`
}

func main() {
	post := Post{
		Id:      "1",
		Content: "Hello World!",
		Author: Author{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}
	// // output, err := xml.Marshal(&post)
	// // 見栄えをよくしたい場合
	// output, err := xml.MarshalIndent(&post, "", "\t")
	// if err != nil {
	// 	fmt.Println("Error marshalling to XML:", err)
	// 	return
	// }
	// // err = ioutil.WriteFile("post.xml", output, 0644)
	// // XML宣言をつける場合
	// err = ioutil.WriteFile("post.xml", []byte(xml.Header + string(output)), 0644)
	// if err != nil {
	// 	fmt.Println("Error writing XML to file:", err)
	// 	return
	// }

	// 自前でエンコードする場合
	xmlFile, err := os.Create("post.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}

	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding XML to file:", err)
		return
	}
}
