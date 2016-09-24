package main

import (
	"net/http"

	"github.com/sj14/web/controller"
)

func main() {
	serve()
}

func serve() {
	http.HandleFunc("/", controller.PanicRecover(controller.IndexHandler))
	http.HandleFunc("/about", controller.PanicRecover(controller.AboutHandler))
	http.HandleFunc("/testpanic", controller.PanicRecover(controller.TestPanicHandler))

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	http.ListenAndServe(":8080", nil)
}
