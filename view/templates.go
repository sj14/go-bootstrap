package view

import "html/template"

//Compile templates on start
var (
	Templates = template.Must(template.ParseFiles(
		"view/header.html",
		"view/footer.html",
		"view/index.html",
		"view/about.html",
		"view/testpanic.html",

		"view/status/404.html",
		"view/status/500_panic.html",
	))
)
