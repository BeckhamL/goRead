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
	Title    string
	Author   string
	Desc     string
	URL      string
	URLImage string
}

var t = template.Must(template.ParseFiles("templates/index.html"))

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

	var articles = []Article{}

	for _, s := range sources.Articles {

		a := Article{
			Title:    s.Title,
			Author:   s.Author,
			Desc:     s.Content,
			URL:      s.URL,
			URLImage: s.URLToImage,
		}

		articles = append(articles, a)

	}

	HTMLerr := t.ExecuteTemplate(w, "index.html", articles)

	if HTMLerr != nil {
		log.Printf("template parsing err:", HTMLerr)
	}

}

func main() {

	server := mux.NewRouter()
	server.HandleFunc("/", runServ)
	server.HandleFunc("/index", index)

	server.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("templates/styles/"))))
	http.ListenAndServe(":8001", server)
}
