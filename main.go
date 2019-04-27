package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var t = template.Must(template.ParseFiles("templates/index.html"))

func runServ(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "access")
}

func index(w http.ResponseWriter, r *http.Request) {

	var articles = topHeadlines()

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
