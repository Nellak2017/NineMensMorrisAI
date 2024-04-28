package main

import "fmt"

// --- Helpers
func updateHalf(position int, half uint32, state int) uint32 {
	bitPosition := uint(position * 2)
	mask := uint32(3) << bitPosition
	return (half & ^mask) | (uint32(state) << bitPosition)
}

// --- Board representation
// Constants representing the possible states of a position on the board
const (
	Empty  = 0 // Represents '00'
	White  = 1 // Represents '01'
	Black  = 2 // Represents '10'
	Unused = 3 // Represents '11'
)

// MorrisBoard represents the board using a pair of 32-bit integers
type MorrisBoard struct {
	firstHalf  uint32 // 32 bits is first 16 of 21 board positions
	secondHalf uint32 // 32 bits is last 5 of 21 board positions
}

// SetPosition sets the state of a position on the board
func (b *MorrisBoard) SetPosition(position int, state int) {
	if position < 16 {
		b.firstHalf = updateHalf(position, b.firstHalf, state)
	} else {
		b.secondHalf = updateHalf(position-16, b.secondHalf, state)
	}
}

// GetPosition returns the state of a position on the board
func (b *MorrisBoard) GetPosition(position int) int {
	if position < 16 {
		return int((b.firstHalf >> uint(position*2)) & 3)
	} else {
		return int((b.secondHalf >> uint((position-16)*2)) & 3)
	}
}

// --- Main function
func main() {
	var board MorrisBoard
	board.SetPosition(0, Empty)  // Set position 0 to 'x'
	board.SetPosition(1, White)  // Set position 1 to 'W'
	board.SetPosition(2, Black)  // Set position 2 to 'B'
	board.SetPosition(3, Unused) // Set position 3 to nothing assigned

	fmt.Println(board.GetPosition(0)) // Get the state of position 0 (White)
	fmt.Println(board.GetPosition(1)) // Get the state of position 1 (Black)
	fmt.Println(board.GetPosition(2)) // Get the state of position 2 (Empty)
	fmt.Println(board.GetPosition(3)) // Get the state of position 3 (Unused)
}
