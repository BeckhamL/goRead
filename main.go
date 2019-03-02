package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/barthr/newsapi"

	"github.com/gorilla/mux"
)

type Article struct {
	Title  string
	Author string
	Desc   string
}

func runServ(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "access")
}

func index(w http.ResponseWriter, r *http.Request) {

	c := newsapi.NewClient("42b47add38084c6c99964fe24dcf6740", newsapi.WithHTTPClient(http.DefaultClient))

	sources, err := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
		Sources: []string{"cnn", "time"},
	})

	if err != nil {
		panic(err)
	}

	for _, s := range sources.Articles {

		a := Article{
			Title:  s.Title,
			Author: s.Author,
			Desc:   s.Content,
		}

		t, HTMLerr := template.ParseFiles("index.html")

		if HTMLerr != nil {
			log.Printf("template parsing err:", HTMLerr)
		}

		HTMLerr = t.Execute(w, a)

		if HTMLerr != nil {
			log.Printf("template executing err:", HTMLerr)
		}

		//fmt.Println(a.Author)
	}
}

func main() {

	server := mux.NewRouter()
	server.HandleFunc("/", runServ)
	server.HandleFunc("/index", index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", server)
}
