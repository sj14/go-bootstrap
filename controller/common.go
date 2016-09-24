package controller

import (
	"log"
	"net/http"

	"github.com/sj14/web/view"
)

//Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	err := view.Templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Recover from a panic and show the error to the user
func PanicRecover(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {

				log.Printf("PANIC RECOVERED: %v\n", r)

				w.WriteHeader(http.StatusInternalServerError)
				m := map[string]interface{}{
					"Title": "Error :-(",
					"Error": r,
				}
				display(w, "panic", &m)
			}
		}()
		f(w, r)
	}
}
