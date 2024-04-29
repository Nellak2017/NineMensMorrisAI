package main

import (
	"fmt"
	"representation"
	"strconv"
)

// --- Main function
func main() {
	fmt.Println("\nMorris board initialized from nothing")
	var board representation.MorrisBoard        // Using empty MorrisBoard
	board.SetPosition(0, representation.Empty)  // Set position 0 to 'x'
	board.SetPosition(1, representation.White)  // Set position 1 to 'W'
	board.SetPosition(2, representation.Black)  // Set position 2 to 'B'
	board.SetPosition(3, representation.Unused) // Set position 3 to nothing assigned

	fmt.Println(board.GetPosition(0)) // Get the state of position 0 (White)
	fmt.Println(board.GetPosition(1)) // Get the state of position 1 (Black)
	fmt.Println(board.GetPosition(2)) // Get the state of position 2 (Empty)
	fmt.Println(board.GetPosition(3)) // Get the state of position 3 (Unused)

	fmt.Println("\n" + board.String())

	fmt.Println("\nMorris board initialized from a string")
	// Create a MorrisBoard from a string representation
	boardStr := "xxxxxxxxxWxWxxxxBxxxW" // White should be able to close mill (20, 11, 1)
	board2 := representation.MorrisBoardFromString(boardStr)
	if board2 == nil {
		fmt.Println("Invalid board string")
		return
	}
	fmt.Println("\n" + board2.String())

	fmt.Println("\n" + strconv.FormatBool(representation.CloseMill(1, board2, representation.White)))

	flippedBoard := board2.InvertColors()
	fmt.Println("\n   board2 again: " + board2.String())
	fmt.Println("board2 inverted: " + flippedBoard.String())

	// Test GenerateRemove
	boardStr2 := "xxxxxxxxxWxWxxxxBxxxW" // Example board string
	board3 := representation.MorrisBoardFromString(boardStr2).InvertColors()
	if board3 == nil {
		fmt.Println("Invalid board string")
		return
	}
	var removedStates []*representation.MorrisBoard
	representation.GenerateRemove(board3, &removedStates)

	fmt.Println("\nBoard used when considering GenerateRemove " + board3.String())
	fmt.Println("\nPossible board states after black piece removal:")
	for _, b := range removedStates {
		fmt.Println(b.String())
	}

	// Test GenerateAdd (moves for opening)
	boardStr3 := "xxxxxxxxxWxWxxxxBxxxW" // White should be able to close mill (20, 11, 1)
	board4 := representation.MorrisBoardFromString(boardStr3)
	if board4 == nil {
		fmt.Println("Invalid board string")
		return
	}

	fmt.Println("GenerateAdd Board: " + board4.String())

	fmt.Println("\nGenerating new board states by adding a white token:")
	newStates := representation.GenerateAdd(board4, representation.White)

	fmt.Println("\nGenerated Board States:")
	for i, state := range newStates {
		fmt.Printf("State %d:\n%s\n", i+1, state.String())
	}

	fmt.Println("\n ------")
	fmt.Println("Board 4: " + board4.String())
	fmt.Println("Static estimate: ", representation.StaticEstimateOpeningNaive(board4))
}
