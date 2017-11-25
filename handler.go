package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func init() {
	htmlTemplate = template.Must(template.New("arc").Parse(defaultTemplateHTML))
}

var htmlTemplate *template.Template

var defaultTemplateHTML = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Choose Your Own Adventure</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{range .Paragraphs}}
			<p>{{.}}</p>
		{{end}}
		<ul>
			{{range .Options}}
				<li><a href="/{{.Arc}}">{{.Text}}</a></li>
			{{end}}
		</ul>
	</body>
</html>`

func StoryHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "intro"
	} else {
		path = path[1:]
	}
	if arc, ok := h.s[path]; ok {
		err := htmlTemplate.Execute(w, arc)
		if err != nil {
			log.Print(err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Story Arc not found!", http.StatusNotFound)
	}
}
