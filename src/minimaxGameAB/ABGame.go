package main

import (
	"fmt"
	"math"
	"os"
	"representation"
	"strconv"
)

// alphaBetaPruningMid function for the minimax algorithm (mid/late game)
func alphaBetaPruningMid(board *representation.MorrisBoard, player int, depth int, alpha int, beta int, maximizingPlayer bool, nodesEvaluated int) (*representation.MorrisBoard, int, int) {
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
		_, evaluated, estimate := alphaBetaPruningMid(move, 3-player, depth-1, alpha, beta, !maximizingPlayer, nodesEvaluated) // Recursively call alphaBeta for the opponent player
		nodesEvaluated = evaluated

		if maximizingPlayer && estimate > bestEstimate {
			if estimate > bestEstimate {
				bestEstimate, bestBoard = estimate, move // Update the best estimate and board for maximizing player
			}
			alpha = int(math.Max(float64(alpha), float64(estimate))) // Update alpha
		} else {
			if estimate < bestEstimate {
				bestEstimate, bestBoard = estimate, move // Update the best estimate and board for minimizing player
			}
			beta = int(math.Min(float64(beta), float64(estimate))) // Update beta
		}

		// Alpha-beta pruning
		if beta <= alpha {
			break // Beta cutoff
		}
	}
	// Defensive check: ensure that bestBoard is not nil
	if bestBoard == nil {
		return &representation.MorrisBoard{}, nodesEvaluated, bestEstimate
	}
	return bestBoard, nodesEvaluated, bestEstimate
}

func MiniMaxMidAB(board *representation.MorrisBoard, player int, depth int) (*representation.MorrisBoard, int, int) {
	bestBoard, evaluated, maxEstimate := alphaBetaPruningMid(board, player, depth, math.MinInt32, math.MaxInt32, true, 0)
	return bestBoard, evaluated, maxEstimate
}

func MiniMaxMidMainAB() error {
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
	bestMove, nodesEvaluated, maxEstimate := MiniMaxMidAB(board, representation.White, depth) // Assume White goes first and White is maximizer

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
