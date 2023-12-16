package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// req: /:day/:part {input}
// res: html page containing answer

func handleForm(w http.ResponseWriter, r *http.Request) AOC {

	day, _ := strconv.Atoi(r.FormValue("dayNumber"))
	file, _, err := r.FormFile("inputFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	var input []byte
	file.Read(input)
	return NewAOC(day, 1, string(input))
}

func handleAns(w http.ResponseWriter, r *http.Request, aoc *AOC) {
	tmple_file := "./static/answer.tmpl"
	tmpl, _ := template.New(tmple_file).ParseFiles(tmple_file)
	tmpl.Execute(w, aoc)
}

func Server() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			// Display the HTML form if it's not a POST request
			tmple_file := "./question.tmpl"
			tmpl, err := template.New(tmple_file).ParseFiles(tmple_file)
			if err != nil {
				http.Error(w, fmt.Sprintf("%v", tmpl), http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, nil)
			return
		}
		w.Write([]byte("hello"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
