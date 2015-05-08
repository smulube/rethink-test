package main

import (
	"encoding/json"
	rethink "github.com/dancannon/gorethink"
	"log"
	"net/http"
)

type Bookmark struct {
	Title string `gorethink:"title"`
	Url   string `gorethink:"url"`
}

var session = InitDb()

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/new", insertBookmark)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("Error: %v", err)
		return
	}
}

func insertBookmark(res http.ResponseWriter, req *http.Request) {
	b := new(Bookmark)
	json.NewDecoder(req.Body).Decode(b)

	_, err := rethink.Table("bookmarks").Insert(b).RunWrite(session)
	if err != nil {
		log.Fatal(err)
		return
	}

	data, _ := json.Marshal("{'bookmark':'saved'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	rows, err := rethink.Table("bookmarks").Run(session)
	if err != nil {
		log.Fatal(err)
	}

	var bookmarks []Bookmark
	_ = rows.All(&bookmarks)

	data, _ := json.Marshal(bookmarks)

	res.Header().Set("Content-Type", "application/json")
	res.Write(data)
}
