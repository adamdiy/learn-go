package main

import (
	"net/http"
	"os"
	"text/template"
)

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("public")))
}

//This pointer is coming from the "text/templates" package
func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templatess"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	//read all files in the templateFolder object we have
	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths,
				basePath + "/" + pathInfo.Name())
		}
	}
	//spread operator
	result.ParseFiles(*templatePaths...)
	return result
}
