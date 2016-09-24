package controller

import "net/http"

func TestPanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("panic")

	m := map[string]interface{}{
		"Title": "TestPanic",
	}
	display(w, "testpanic", &m)
}
