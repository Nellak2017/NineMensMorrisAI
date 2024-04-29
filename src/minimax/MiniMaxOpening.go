package main

import (
	"fmt"
	"math"
	"os"
	"representation"
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

// alphaBeta function for the minimax algorithm
func alphaBeta(board *representation.MorrisBoard, player int, depth int, maximizingPlayer bool, nodesEvaluated int) (*representation.MorrisBoard, int, int) {
	// Base case: return the static evaluation if we've reached the depth limit
	if depth == 0 {
		nodesEvaluated++
		return board, nodesEvaluated, representation.StaticEstimateOpeningNaive(board)
	}

	var bestBoard *representation.MorrisBoard
	var bestEstimate int

	if maximizingPlayer {
		bestEstimate = math.MinInt32
	} else {
		bestEstimate = math.MaxInt32
	}

	moves := representation.GenerateAdd(board, player) // Generate possible moves for the current board state
	// Defensive check: ensure that moves is not nil
	if moves == nil {
		return &representation.MorrisBoard{}, nodesEvaluated, bestEstimate
	}

	for _, move := range moves {
		_, evaluated, estimate := alphaBeta(move, 3-player, depth-1, !maximizingPlayer, nodesEvaluated) // Recursively call alphaBeta for the opponent player
		nodesEvaluated = evaluated

		if (maximizingPlayer && estimate > bestEstimate) || (!maximizingPlayer && estimate < bestEstimate) {
			bestEstimate, bestBoard = estimate, move // Update the best estimate and board based on maximizing or minimizing player
		}
	}
	// Defensive check: ensure that bestBoard is not nil
	if bestBoard == nil {
		return &representation.MorrisBoard{}, nodesEvaluated, bestEstimate
	}
	return bestBoard, nodesEvaluated, bestEstimate
}

func MiniMax(board *representation.MorrisBoard, player int, depth int) (*representation.MorrisBoard, int, int) {
	bestBoard, evaluated, maxEstimate := alphaBeta(board, player, depth, true, 0)
	return bestBoard, evaluated, maxEstimate
}

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
	// fmt.Printf("Input file: %s, Output file: %s, Depth: %d\n", file1, file2, depth)

	// Read input board file
	inputBoard, err := os.ReadFile(file1)
	if err != nil {
		return fmt.Errorf("failed to read input board file: %v", err)
	}

	// Print input board position
	fmt.Printf("Input position: %s\n", inputBoard)

	// Convert inputBoard from string to board using the method, MorrisBoardFromString
	board := representation.MorrisBoardFromString(string(inputBoard))
	if board == nil {
		return fmt.Errorf("invalid input board")
	}
	// Compute min-max algorithm values
	bestMove, nodesEvaluated, maxEstimate := MiniMax(board, representation.White, depth) // Assume White goes first

	// Print output, positions evaluated, minimax estimate
	fmt.Printf("Output position: %s\n", bestMove.String())
	fmt.Printf("Positions evaluated by static estimation: %d\n", nodesEvaluated)
	fmt.Printf("MINIMAX estimate: %d\n", maxEstimate)

	// Write output board to output file
	output := []byte(bestMove.String())
	if err := os.WriteFile(file2, output, 0644); err != nil {
		return fmt.Errorf("failed to write output board file: %v", err)
	}

	return nil
}
