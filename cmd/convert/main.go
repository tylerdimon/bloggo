package main

import (
	"blog"
)

func main() {
	inputFolder := "./posts"
	outputFolder := "./html"
	if err := blog.Convert(inputFolder, outputFolder); err != nil {
		panic(err)
	}
}
