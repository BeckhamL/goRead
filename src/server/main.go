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

	var articles = []Article{}

	for _, s := range sources.Articles {

		a := Article{
			Title:  s.Title,
			Author: s.Author,
			Desc:   s.Content,
		}

		articles = append(articles, a)

	}

	html := `<!DOCTYPE html>
	<html>
	<body>

	{{range $articles := .}}
			<p> {{$articles.Title}}</p>
			<p> {{$articles.Author}}</p>
			<p> {{$articles.Desc}}</p>
	{{end}}

	</body>
	</html>`

	t, HTMLerr := template.New("test").Parse(string(html))

	if HTMLerr != nil {
		log.Printf("template parsing err:", HTMLerr)
	}

	HTMLerr = t.Execute(w, articles)

}

func main() {

	server := mux.NewRouter()
	server.HandleFunc("/", runServ)
	server.HandleFunc("/index", index)
	http.ListenAndServe(":8001", server)
}
