package main

import (
	"blog"
	"net/http"
)

func main() {
	blog.GetConfig()

	fileServer := http.FileServer(http.Dir("html"))

	http.Handle("/blog/", http.StripPrefix("/blog/", fileServer))

	blog.ServeCustomPage("resume")
	blog.ServeCustomPage("")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
