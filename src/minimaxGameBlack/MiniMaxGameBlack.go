package main

import (
	"fmt"
	"math"
	"os"
	"representation"
	"strconv"
)

// alphaBetaMid function for the minimax algorithm (mid/late game)
func alphaBetaMid(board *representation.MorrisBoard, player int, depth int, maximizingPlayer bool, nodesEvaluated int) (*representation.MorrisBoard, int, int) {
	// Base case: return the static evaluation if we've reached the depth limit
	if depth == 0 {
		nodesEvaluated++
		return board, nodesEvaluated, representation.StaticEstimateMidgameEndgame(board)
	}

	var bestBoard *representation.MorrisBoard
	var bestEstimate int

	if maximizingPlayer {
		bestEstimate = math.MinInt32
	} else {
		bestEstimate = math.MaxInt32
	}

	moves := representation.GenerateMovesMidgameEndgame(board, player) // Generate possible moves for the current board state

	// Defensive check: ensure that moves is not nil
	if moves == nil {
		return &representation.MorrisBoard{}, nodesEvaluated, bestEstimate
	}

	for _, move := range moves {
		_, evaluated, estimate := alphaBetaMid(move, 3-player, depth-1, !maximizingPlayer, nodesEvaluated) // Recursively call alphaBetaMid for the opponent player
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

func MiniMaxMid(board *representation.MorrisBoard, player int, depth int) (*representation.MorrisBoard, int, int) {
	bestBoard, evaluated, maxEstimate := alphaBetaMid(board, player, depth, false, 0) // assuming black is minizing player
	return bestBoard, evaluated, maxEstimate
}

func MiniMaxMidMain() error {
	// Check number of command-line arguments, early return if invalid
	if len(os.Args) < 4 {
		return fmt.Errorf("usage: MiniMaxMid <input_file> <output_file> <depth>")
	}

	// Extract command-line arguments
	file1, file2, depthStr := os.Args[1], os.Args[2], os.Args[3]

	// Convert depth string to integer
	depth, err := strconv.Atoi(depthStr)
	if err != nil {
		return fmt.Errorf("invalid depth: %s", depthStr)
	}

	// Process the command-line arguments (Debugging only)
	//fmt.Printf("Input file: %s, Output file: %s, Depth: %d\n", file1, file2, depth)

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
	bestMove, nodesEvaluated, maxEstimate := MiniMaxMid(board, representation.Black, depth) // Assume White goes first and White is maximizer

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
