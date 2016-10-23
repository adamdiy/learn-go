package main

import (
	"net/http"
	"os"
	"text/template"
	"bufio"
	"strings"
)

func main() {
	templates := populateTemplates()
	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {
		requestedFile := req.URL.Path[1:]
		template := 
			templates.Lookup(requestedFile + ".html")

			if template != nil {
				template.Execute(w, nil)
			} else {
				w.WriteHeader(404)
			}
		})
		http.ListenAndServe(":8000", nil)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plan"
	}
}

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath + "/" + pathInfo.Name())
		}
	}
	result.ParseFiles(*templatePaths...)
	return result
}