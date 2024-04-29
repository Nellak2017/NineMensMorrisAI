package main

import (
	"fmt"
	//"representation"
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
		newStates := representation.GenerateAdd(board, player)
		for i, state := range newStates {
			fmt.Printf("State %d:\n%s\n", i+1, state.String())
		}

			bestMove, nodesEvaluated, maxEstimate := MiniMax(board, player, 1) // Call MiniMax to find the best move for the current player
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
