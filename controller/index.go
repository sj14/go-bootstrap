package controller

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		display(w, "404", nil)
		return
	}
	m := map[string]interface{}{
		"Title": "Home",
	}
	display(w, "index", &m)
}
