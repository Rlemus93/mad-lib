package main

import (
	"log"
	"net/http"

	"github.com/rlemus93/multi-page-form/web"
)

type Form struct {
	Type  string //noun, verb, adj - custom attr "data-pos"
	Name  string //id
	Value string //data
}

type MadLib struct {
	Nouns    []string
	Adj      []string
	Verb     []string
	Pronouns []string
}

func HandlerIndex(w http.ResponseWriter, r *http.Request) {

	web.ExeTemp(w, "index", nil)
}

func HandlerParseForm(w http.ResponseWriter, r *http.Request) {

	// r.ParseForm()
	// var nouns []string
	// nouns[0] = r.FormValue("noun1")
	// //rwad req

	// ml := MadLib{
	// 	Nouns: nouns,
	// }

	// web.ExeTemp(w, "tacostand", ml)
}

func main() {

	web.LoadTemplates()

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public", fs))

	http.HandleFunc("/", HandlerIndex)
	http.HandleFunc("/submit", HandlerParseForm)

	// http.ListenAndServe(":8080", nil)
	println("http://0.0.0.0:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
