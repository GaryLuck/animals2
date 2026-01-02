package main

import (
	"encoding/json"
	"os"
)

const defaultFilename = "database.json"

// Save writes the tree to a JSON file.
func Save(root *Node, filename string) error {
	data, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// Load reads the tree from a JSON file.
// If the file doesn't exist, it returns a default starting tree.
func Load(filename string) (*Node, error) {
	data, err := os.ReadFile(filename)
	if os.IsNotExist(err) {
		// Return default tree
		return &Node{Text: "Human"}, nil
	}
	if err != nil {
		return nil, err
	}

	var root Node
	if err := json.Unmarshal(data, &root); err != nil {
		return nil, err
	}
	return &root, nil
}
