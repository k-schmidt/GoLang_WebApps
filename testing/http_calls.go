package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Println(t)
}

// main ...
func main() {
	routes := mux.NewRouter()
	routes.HandleFunc("/test", testHandler)
	http.ListenAndServe(":8080", nil)
}
