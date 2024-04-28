package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Requirements:

 1. get three command line args
    X file 1 (board 1, the input file name)
    X file 2 (board 2, the output file name)
    X depth of tree to be searched

2. output:
  - input board position (list of 21 characters) before WHITE plays its best move
  - command line
  - output board position (list of 21 characters) after WHITE plays its best move as determined by minimax search tree
  - command line
  - output file
  - number of positions evaluated by the static estimation function
  - command line
  - min-max estimate for that move
  - command line

3. minimax search tree uses:
  - depth given by command line argument
  - static estimation function given in Morris-B.pdf

Additional considerations:

	Use the move generator and the static estimation function for the opening phase.
	Don't verify that the position is an opening position.
	Assume that this game never goes into the midgame phase.
*/

/*

func minimax(board, player = White) {

}

*/

func MiniMaxOpeningMain() error {
	// Check number of command-line arguments, early return if invalid
	if len(os.Args) < 4 {
		return fmt.Errorf("usage: MiniMaxOpening <input_file> <output_file> <depth>")
	}

	// Extract command-line arguments
	file1, file2, depthStr := os.Args[1], os.Args[2], os.Args[3]

	// Convert depth string to integer
	depth, err := strconv.Atoi(depthStr)
	if err != nil {
		return fmt.Errorf("invalid depth: %s", depthStr)
	}

	// Process the command-line arguments (Debugging only)
	fmt.Printf("Input file: %s, Output file: %s, Depth: %d\n", file1, file2, depth)

	// Read input board file
	inputBoard, err := os.ReadFile(file1)
	if err != nil {
		return fmt.Errorf("failed to read input board file: %v", err)
	}

	// Print input board position
	fmt.Printf("Input board position: %s\n", inputBoard)

	// Perform the actual processing here (e.g., reading input file, performing minimax algorithm, writing output file)

	// Print output, positions evaluated, minimax estimate
	fmt.Printf("Output position: %s\n", "no output yet")
	fmt.Printf("Positions evaluated by static estimation: %s\n", "no positions evaluated yet")
	fmt.Printf("MINIMAX estimate: %s\n", "no estimate yet")

	return nil
}
