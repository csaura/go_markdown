package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/russross/blackfriday"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome To this markdown parser!\n")
}

func GetMarkdown(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var markdown_text string

	params := r.URL.Query()
	markdown_text = ConvertTextToMarkdown(params)

	fmt.Fprintf(w, markdown_text)
}

func PostMarkdown(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var markdown_text string

	r.ParseForm()
	params := r.PostForm

	markdown_text = ConvertTextToMarkdown(params)

	fmt.Fprintf(w, markdown_text)
}

func ConvertTextToMarkdown(m map[string][]string) string {
	var markdown_text string
	text, exists := m["text"]

	if exists {
		fmt.Println(text[0])
		output := blackfriday.MarkdownCommon([]byte(text[0]))
		markdown_text = string(output)
		fmt.Println(markdown_text)
	} else {
		fmt.Println("no text")
		markdown_text = "no text"
	}

	return markdown_text
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/markdown", GetMarkdown)
	router.POST("/markdown", PostMarkdown)

	log.Fatal(http.ListenAndServe(":8080", router))
}
