package blog

import (
	"net/http"
)

const HtmlDir = "html/"
const HomePage = ""
const HomePageHtmlFile = "index.html"

// ServeCustomPage passing in a name will serve a file with that name at a path with that name
// resume would serve html/resume.html at /resume/
func ServeCustomPage(name string) {
	var route string
	var filepath string
	if name == HomePage {
		route = "/"
		filepath = HtmlDir + HomePageHtmlFile
	} else {
		route = "/" + name + "/"
		filepath = HtmlDir + name + ".html"
	}

	http.Handle(route, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath)
	}))
}
