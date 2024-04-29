package main

import (
	"fmt"
)

func main() {
	bestMove := MiniMaxOpeningMain()
	fmt.Printf("%v", bestMove)
	/*
		// --- Test min-max function
		board, player := representation.MorrisBoardFromString("xxxxxxWxxBxxxxxxBxWxx"), representation.White
		if board == nil {
			fmt.Printf("Your board is undefined")
			return
		}
		bestMove, nodesEvaluated, maxEstimate := MiniMaxAB(board, player, 1) // Call MiniMax to find the best move for the current player
		fmt.Println("Best Move:")
		if bestMove != nil {
			fmt.Println(bestMove.String()) // Display the best move's board state as a string
		} else {
			fmt.Println("No valid move found.")
		}
		fmt.Println("Nodes Evaluated:", nodesEvaluated)
		fmt.Println("Max Estimate:", maxEstimate)
	*/
}
