package main

import (
	"fmt"
	"os"
	"uncut/internal/app/gui"
)

func loadFile(filePath string) (content []byte) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Errorf("failed to load file %s [%w]", filePath, err))
	}
	return content
}

func GenerateLead() {
	// Load lead layout from json files
	schedule := loadFile("assets/schedule.json")
	content := loadFile("assets/content.json")

	app := gui.NewApp("")

	s := "{\"strategy\": \"lightestBin\"}"

	app.GenerateLeads(string(schedule), string(content), []int{1}, "/Users/luishankel/Downloads/", s)
}
