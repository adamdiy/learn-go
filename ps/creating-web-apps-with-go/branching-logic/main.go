package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		tmpl, err := template.New("test").Parse(doc)
		if err == nil {
			//context := Context{"the message"}
			//tmpl.Execute(w, context)
			tmpl.Execute(w, req.URL.Path)
		}
	})

	http.ListenAndServe(":8080", nil)

}

const doc = `
<!DOCTYPE html>
<html>
	<head><title>Example Title</title></head>
	<body>
		{{if eq . "/Google"}}
			<h1>Hey, Google made Go</h1>
		{{else}}
			<h1>Hello, {{.}}<h1>
		{{end}}
	</body>
</htm>
`

type Context struct {
	Message string
}
