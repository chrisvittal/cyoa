package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var defaultTemplateCLI = `
{{.Title}}:
{{range .Paragraphs}}{{printf "  %s\n" .}}{{end}}
{{range .Options}}{{printf "%12s: %s\n" .Arc .Text}}{{end}}
`

var cliTemplate *template.Template

func init() {
	cliTemplate = template.Must(template.New("arc").Parse(defaultTemplateCLI))
}

func StoryCLI(s Story) {
	next := "intro"

	for arc, ok := s[next]; ok; arc, ok = s[next] {
		err := cliTemplate.Execute(os.Stdout, arc)
		if err != nil {
			log.Fatal(err)
		}
		if len(arc.Options) == 0 {
			break
		}
		fmt.Print("Where to next? ")
		fmt.Scanf("%s", &next)
	}
}
