package main

import (
	"net/http"
	"os"
	"text/template"
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

	http.ListenAndServe(":8080", nil)

}

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatesPathRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)
	for _, pathInfo := range templatesPathRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths,
			basePath + "/" + pathInfo.Name())
		}
	}
	result.ParseFiles(*templatePaths...)
	return result
}
