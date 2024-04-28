package main

import (
	"NineMensMorrisAI/minimax"
	"fmt"
)

func main() {
	bestMove := minimax.MiniMaxOpeningMain()

	fmt.Printf("%v", bestMove)
}
