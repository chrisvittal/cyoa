package main

import (
	"encoding/json"
	"io"
)

// A Story is a collection of StoryArcs indexed by name.
type Story map[string]StoryArc

// StoryArc represents a single arc of a story.
type StoryArc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option represents a choice the reader can make for which StoryArc to follow
// next.
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// ParseStory takes json and parses it into a Story.
func ParseStory(jsn io.Reader) (Story, error) {
	var story Story
	decoder := json.NewDecoder(jsn)
	err := decoder.Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil
}
