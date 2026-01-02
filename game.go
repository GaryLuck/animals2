package main

import (
	"bufio"
	"fmt"
	"strings"
)

// Node represents a node in the decision tree.
// It can be a question (internal node) or an animal (leaf node).
type Node struct {
	Text string `json:"text"`          // The question or the animal name
	Yes  *Node  `json:"yes,omitempty"` // Child node for 'yes' answer
	No   *Node  `json:"no,omitempty"`  // Child node for 'no' answer
}

// IsLeaf checks if the node is a leaf node (an animal).
func (n *Node) IsLeaf() bool {
	return n.Yes == nil && n.No == nil
}

// Play handles a single round of the game starting from the current node.
// It returns true if the game logic modified the tree (learned something), false otherwise.
func Play(n *Node, reader *bufio.Reader) bool {
	if n.IsLeaf() {
		return handleLeaf(n, reader)
	}

	// It's a question node
	if askYesNo(n.Text, reader) {
		return Play(n.Yes, reader)
	} else {
		return Play(n.No, reader)
	}
}

// handleLeaf handles the guess when we reach a leaf node.
func handleLeaf(n *Node, reader *bufio.Reader) bool {
	if askYesNo("Is it a "+n.Text+"?", reader) {
		fmt.Println("I win!")
		return false
	}

	// Computer lost, time to learn
	fmt.Println("I lost! What was the animal you were thinking of?")
	newAnimalName := readLine(reader)

	fmt.Println("Please type a yes/no question that distinguishes a " + newAnimalName + " from a " + n.Text + ".")
	question := readLine(reader)

	fmt.Println("As a " + newAnimalName + ", " + question + " (yes/no)?")
	isYesForNew := askYesNo("", reader)

	// Create new nodes
	// The current node 'n' becomes the question node.
	// We need to back up the old animal (n.Text) to a new child node.
	oldAnimal := &Node{Text: n.Text}
	newAnimal := &Node{Text: newAnimalName}

	n.Text = question
	if isYesForNew {
		n.Yes = newAnimal
		n.No = oldAnimal
	} else {
		n.Yes = oldAnimal
		n.No = newAnimal
	}

	return true
}

// Helper functions for input

func askYesNo(question string, reader *bufio.Reader) bool {
	if question != "" {
		fmt.Println(question + " (yes/no)")
	}
	for {
		input := strings.ToLower(readLine(reader))
		if input == "y" || input == "yes" {
			return true
		}
		if input == "n" || input == "no" {
			return false
		}
		fmt.Println("Please answer yes or no.")
	}
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}
