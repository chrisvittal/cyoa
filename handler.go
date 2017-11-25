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
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      <ul>
        {{range .Options}}
          <li><a href="/{{.Arc}}">{{.Text}}</a></li>
        {{end}}
      </ul>
    </section>
  </body>
  <style>
    body {
      font-family: helvetica, arial;
    }
    h1 {
      text-align:center;
      position:relative;
    }
    .page {
      width: 80%;
      max-width: 500px;
      margin: auto;
      margin-top: 40px;
      margin-bottom: 40px;
      padding: 80px;
      background: #FFFCF6;
      border: 1px solid #eee;
      box-shadow: 0 10px 6px -6px #777;
    }
    ul {
      border-top: 1px dotted #ccc;
      padding: 10px 0 0 0;
      -webkit-padding-start: 0;
    }
    li {
      padding-top: 10px;
    }
    a,
    a:visited {
      text-decoration: none;
      color: #6295b5;
    }
    a:active,
    a:hover {
      color: #7792a2;
    }
    p {
      text-indent: 1em;
    }
  </style>
</html>`

// HandlerOption is used to set options for our http.Handler we use to handle
// going through the story in the webapp.
type HandlerOption func(h *handler)

// WithTemplate is used to set the template for the handler.
func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

// WithPathFunc is used to set the function that parses the paths in order to
// look up parts of the Story
func WithPathFunc(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFunc = fn
	}
}

// StoryHandler creates a new story with the appropriate options set.
func StoryHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, htmlTemplate, defaultPathFunc}
	for _, opt := range opts {
		opt(&h)
	}

	return h
}

type handler struct {
	s        Story
	t        *template.Template
	pathFunc func(r *http.Request) string
}

func defaultPathFunc(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "intro"
	} else {
		path = path[1:]
	}
	return path
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFunc(r)
	if arc, ok := h.s[path]; ok {
		err := h.t.Execute(w, arc)
		if err != nil {
			log.Print(err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Story Arc not found!", http.StatusNotFound)
	}
}
