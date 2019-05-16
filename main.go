package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var t = template.Must(template.ParseFiles("templates/index.html"))

func runServ(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "templates/intro.html")
}

// CNN route
func CNN(w http.ResponseWriter, r *http.Request) {

	var articles = topHeadlinesCNN()

	HTMLerr := t.ExecuteTemplate(w, "index.html", articles)

	if HTMLerr != nil {
		log.Printf("template parsing err:", HTMLerr)
	}
}

// BBC route
func BBCNews(w http.ResponseWriter, r *http.Request) {

	var articles = topHeadlinesBBCNews()

	HTMLerr := t.ExecuteTemplate(w, "index.html", articles)

	if HTMLerr != nil {
		log.Printf("template parsing err:", HTMLerr)
	}
}

func main() {

	server := mux.NewRouter()
	server.HandleFunc("/", runServ)
	server.HandleFunc("/CNN", CNN)
	server.HandleFunc("/BBCNews", BBCNews)

	server.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("templates/styles/"))))
	http.ListenAndServe(":8001", server)
}
