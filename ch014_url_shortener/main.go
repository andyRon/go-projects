package main

import (
	"crypto/md5"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// TODO

var urlStore = make(map[string]string)

func shorten(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.Form.Get("url")
	shortURL := fmt.Sprintf("%x", md5.Sum([]byte(url)))[:5]
	urlStore[shortURL] = url
	fmt.Fprintf(w, "http://localhost:8085/%s\n", shortURL)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create", shorten).Methods("POST")
	r.HandleFunc("/{shortURL}", redirect).Methods("GET")
	log.Fatal(http.ListenAndServe("8085", r))

}

func redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	originalURL, ok := urlStore[vars["shortURL"]]
	if !ok {
		http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
	} else {
		http.NotFound(w, r)
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

/*
curl -X POST http://localhost:8085/create -d "url=http://google.com"
*/
