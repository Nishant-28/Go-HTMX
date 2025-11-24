package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title, Director string
}

func main() {
	fmt.Println("Go HTMX App...")
	h1 := func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello World")
		// io.WriteString(w, r.Method)

		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis"},
				{Title: "Pulp Fiction", Director: "Quentin"},
				{Title: "Inception", Director: "Christopher"},
			},
		}
		tmpl.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		// log.Print("HTMX request received")
		// log.Print(r.Header.Get("HX-Request"))
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		// htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
		// tmpl, _ := template.New("t").Parse(htmlStr)
		// tmpl.Execute(w, nil)
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})

	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
