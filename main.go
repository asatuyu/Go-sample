package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main () {
	jsonFile, err := os.Open("sample.json")
	if err != nil {
		fmt.Println("Jsonファイルが開けません", err)
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("JSONデータを読み込めません", err)
		return
	}

	var post Post
	json.Unmarshal(jsonData, &post)

	fmt.Println(post)
	fmt.Println(post.Comments)
	fmt.Println(post.Comments[0].Content)
}