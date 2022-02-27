package main

import (
	"net/http"
	"strings"
)

func f1(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Set(k, strings.Join(v, ""))
	}
}

func main() {
	http.HandleFunc("/assignment/httpServer/", f1)
	http.ListenAndServe("127.0.0.1:7070", nil)
}
