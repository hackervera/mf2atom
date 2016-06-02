package main

import (
	"net/http"

	"github.com/tjgillies/mf2atom"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.Form.Get("url")
	atom := mf2atom.Parse(url)
	w.Write([]byte(atom))
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
