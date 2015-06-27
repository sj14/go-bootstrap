package main

import (
	"html/template"
	"net/http"
	"time"
)

//Compile templates on start
var (
	templates = template.Must(template.ParseFiles(
		"html/header.html",
		"html/footer.html",
		"html/index.html",
		"html/about.html",
	))
	startTime time.Time
)

//Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//The handlers.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]interface{})
	m["Title"] = "Home"
	m["Uptime"] = time.Now().Sub(startTime)

	display(w, "index", &m)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]interface{})
	m["Title"] = "About"

	display(w, "about", &m)
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	startTime = time.Now()

	//Listen on port 8080
	http.ListenAndServe(":8080", nil)
}
