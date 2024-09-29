package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Router(port int) {
	router := mux.NewRouter()
	router.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "The %s page is: %s", page, title)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		return
	}

}
