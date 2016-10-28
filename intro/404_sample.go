package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
)

// pageHandler ...
func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	fileName := "files/" + pageID + ".html"
	if err, _ := os.Stat(fileName); err == nil {
		fileName = "static/404.html"
	}
	http.ServeFile(w, r, fileName)
}

// main ...
func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	rtr.HandleFunc("/homepage", pageHandler)
	rtr.HandleFunc("/contact", pageHandler)
	http.Handle("/", rtr)
	http.ListenAndServe(PORT, nil)
}
