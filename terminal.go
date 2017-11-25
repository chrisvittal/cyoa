package main

import (
	"fmt"
	"go/doc"
	"os"
)

// StoryCLI is the entry point to the terminal version of the application
func StoryCLI(s Story) {
	next := "intro"

	for arc, ok := s[next]; ok; arc, ok = s[next] {
		printArc(arc)
		if len(arc.Options) == 0 {
			println("    The End.")
			break
		} else {
			fmt.Print("Where to next? ")
			fmt.Scanf("%s", &next)
			println()
		}
	}
}

func printArc(arc StoryArc) {
	fmt.Printf("%s:\n", arc.Title)
	for _, p := range arc.Paragraphs {
		// This is a hack to allow indents for each paragraph
		print("    ")
		doc.ToText(os.Stdout, p, "", "", 80)
	}
	for _, o := range arc.Options {
		fmt.Printf("%12s: %s\n", o.Arc, o.Text)
	}
}
