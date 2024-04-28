package main

import (
	"os"
	"testing"
)

// Test that MiniMaxOpening allows for valid command line arguments
func TestCommandLineArguments(t *testing.T) {
	// Save original os.Args and defer restoring it
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Simulate command-line arguments
	os.Args = []string{"main.go", "board1.txt", "board2.txt", "3"}

	// Call your main program function with the simulated command-line arguments
	err := MiniMaxOpeningMain()

	// Check if any errors occurred during program execution
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
