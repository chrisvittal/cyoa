package main

import (
	"encoding/json"
	"io"
)

// A whole story, consisting of many arcs
type Story map[string]StoryArc

// A story arc, complete with title, text, and follow up choices
type StoryArc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Choices for the end of an arc
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// Takes a Reader (assumed to be JSON) and parses it into a Story
func ParseStory(jsn io.Reader) (Story, error) {
	var story Story
	decoder := json.NewDecoder(jsn)
	err := decoder.Decode(&story)
	if err != nil {
		return nil, err
	} else {
		return story, nil
	}
}
