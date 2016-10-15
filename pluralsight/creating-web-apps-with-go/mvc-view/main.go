package main

import (
	"net/http"
	"os"
	"text/template"
)

func main() {
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		requestedFile := req.URL.Path[1:]
		template := templates.Lookup(requestedFile + ".html")

		if template != nil {
			template.Execute(w, nil)
		} else {
			w.WriteHeader(404)
		}
	})
	http.ListenAndServe(":8080", nil)
}


//This pointer is coming from the "text/templates" package
func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
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
