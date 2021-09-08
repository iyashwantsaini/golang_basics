package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
	"webbasics/rps"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	rendertemplate(w, "index.html")
}

func playround(w http.ResponseWriter, r *http.Request) {
	result := rps.PlayRound(1)
	out, err := json.MarshalIndent(result, "", "	")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {
	http.HandleFunc("/play", playround)
	http.HandleFunc("/", homepage)
	log.Println("starting web server on post 8080")
	http.ListenAndServe(":8080", nil)
}

func rendertemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	//check error while parsing
	if err != nil {
		log.Println(err)
	}
	//check error while execution
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
