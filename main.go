package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	isatty "github.com/mattn/go-isatty"
)

var fileName string
var cmdLine bool
var port int

func init() {
	flag.BoolVar(&cmdLine, "text", false, "run the application as a commandline app, outputting to the terminal")
	flag.StringVar(&fileName, "file", "gopher.json", "JSON file containing the story data")
	flag.IntVar(&port, "port", 3000, "port to start the CYOA webapp on")
	flag.Parse()
}

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	story, err := ParseStory(file)
	if err != nil {
		log.Fatal(err)
	}

	runServer(story)
}

func runServer(story Story) {
	h := StoryHandler(story)
	log.Print("Starting the server on port:", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), h))
}
