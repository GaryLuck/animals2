package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Load database
	root, err := Load(defaultFilename)
	if err != nil {
		fmt.Printf("Error loading database: %v\n", err)
		return
	}

	fmt.Println("Welcome to the Animals Game!")

	for {
		changed := Play(root, reader)
		if changed {
			// Save immediately if learned something, or we could save at exit.
			// Saving immediately prevents data loss on crash/force close.
			if err := Save(root, defaultFilename); err != nil {
				fmt.Printf("Error saving database: %v\n", err)
			}
		}

		fmt.Println("Do you want to play again? (yes/no)")
		input := strings.ToLower(readLine(reader))
		if input != "y" && input != "yes" {
			break
		}
	}

	// Final save
	if err := Save(root, defaultFilename); err != nil {
		fmt.Printf("Error saving database: %v\n", err)
	}
	fmt.Println("Goodbye!")
}
