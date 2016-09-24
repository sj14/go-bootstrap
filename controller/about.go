package controller

import "net/http"

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{
		"Title": "About",
	}
	display(w, "about", &m)
}
