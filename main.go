package main

import (
	"html/template"
	"net/http"
)

//Compile templates on start
var templates = template.Must(template.ParseFiles(
	"html/header.html",
	"html/footer.html",
	"html/index.html",
	"html/about.html",
))

//A Page structure
type Page struct {
	Title string
}

//Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

//The handlers.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "index", &Page{Title: "Home"})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "about", &Page{Title: "About"})
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	//Listen on port 8080
	http.ListenAndServe(":8080", nil)
}
