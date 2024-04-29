package representation

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
	}
	return int((b.secondHalf >> uint((position-16)*2)) & 3)
}

// MorrisBoardFromString creates a MorrisBoard from a string representation
func MorrisBoardFromString(boardStr string) *MorrisBoard {
	if len(boardStr) != 21 {
		return nil // Board string must have exactly 21 characters
	}

	board := &MorrisBoard{}

	for i, char := range boardStr {
		switch char {
		case 'x':
			board.SetPosition(i, Empty)
		case 'W':
			board.SetPosition(i, White)
		case 'B':
			board.SetPosition(i, Black)
		default:
			return nil // Invalid character in board string
		}
	}

	return board
}

// String returns a string representation of the MorrisBoard state
func (b *MorrisBoard) String() string {
	var result string
	for i := 0; i < 21; i++ {
		switch b.GetPosition(i) {
		case Empty:
			result += "x"
		case White:
			result += "W"
		case Black:
			result += "B"
		case Unused:
			result += "-"
		}
	}
	return result
}

// InvertColors flips the colors on the board (White becomes Black and Black becomes White)
func (b *MorrisBoard) InvertColors() *MorrisBoard {
	flippedBoard := &MorrisBoard{}

	for i := 0; i < 21; i++ {
		currentState := b.GetPosition(i)
		switch currentState {
		case White:
			flippedBoard.SetPosition(i, Black)
		case Black:
			flippedBoard.SetPosition(i, White)
		default:
			flippedBoard.SetPosition(i, currentState)
		}
	}

	return flippedBoard
}

// CloseMill checks if placing a token at position j closes a mill on the board b
func CloseMill(j int, b *MorrisBoard, color int) bool {
	C := b.GetPosition(j) // Color of the token placed at position j (0, 1, 2)

	if C != 0 {
		return false // if it is not empty, return false
	}

	// Define mill formations as arrays of positions
	mills := [][]int{
		{0, 6, 18}, {0, 2, 4}, //   0 === a0
		{1, 11, 20},           //   1 === g0
		{2, 0, 4}, {2, 7, 15}, //   2 === b1
		{3, 10, 17},           //   3 === f1
		{4, 2, 0}, {4, 8, 12}, //   4 === c2
		{5, 9, 14},            //   5 === e2
		{6, 0, 18}, {6, 7, 8}, //   6 === a3
		{7, 2, 15}, {7, 6, 8}, //   7 === b3
		{8, 4, 12}, {8, 6, 7}, //   8 === c3
		{9, 5, 14}, {9, 10, 11}, //   9 === e3
		{10, 3, 17}, {10, 9, 11}, // 10 === f3
		{11, 1, 20}, {11, 9, 10}, // 11 === g3
		{12, 4, 8}, {12, 13, 14}, // 12 === c4
		{13, 12, 14}, {13, 16, 19}, // 13 === d4
		{14, 12, 13}, {14, 5, 9}, // 14 === e4
		{15, 7, 2}, {15, 16, 17}, // 15 === b5
		{16, 13, 19}, {16, 15, 17}, // 16 === d5
		{17, 3, 10}, {17, 15, 16}, // 17 === f5
		{18, 0, 6}, {18, 19, 20}, // 18 === a6
		{19, 13, 16}, {19, 18, 20}, // 19 === d6
		{20, 11, 1}, {20, 18, 19}, // 20 === g6
	}

	// Iterate through the defined mills and check if the position j completes any of them
	for _, mill := range mills {
		var mil0, mil1, mil2 = mill[0], mill[1], mill[2]
		var pos1, pos2 = b.GetPosition(mil1), b.GetPosition(mil2)
		if j == mil0 && pos1 == color && pos2 == color {
			return true
		}
	}

	return false
}

func GenerateRemove(board *MorrisBoard, L *[]*MorrisBoard) {
	// Check if removing each black piece does not close a mill, since normally you can only remove a black piece if it is not a mill
	for location := 0; location < 21; location++ {
		if board.GetPosition(location) == Black {
			b := &MorrisBoard{firstHalf: board.firstHalf, secondHalf: board.secondHalf} // Create a copy of the board
			b.SetPosition(location, Empty)                                              // Remove the black piece from the specified location
			if !CloseMill(location, b, Black) {                                         // Check if removing this piece does not close a mill
				*L = append(*L, b) // If it doesn't close a mill, add the resulting board state to the list L
			}
		}
	}

	// If no states were added (all black pieces are in mills), add all possible removal states.
	// This is because if all pieces are mills then you still remove one
	if len(*L) == 0 {
		for location := 0; location < 21; location++ {
			if board.GetPosition(location) == Black { // Check if the position contains a black piece
				b := &MorrisBoard{firstHalf: board.firstHalf, secondHalf: board.secondHalf} // Create a copy of the board
				b.SetPosition(location, Empty)                                              // Remove the black piece from the specified location
				*L = append(*L, b)                                                          // Add the resulting board state to the list L
			}
		}
	}
}

func GenerateAdd(board *MorrisBoard, color int) []*MorrisBoard {
	L := []*MorrisBoard{}

	for location := 0; location < 21; location++ {
		if board.GetPosition(location) == Empty {
			b := &MorrisBoard{firstHalf: board.firstHalf, secondHalf: board.secondHalf} // Create a copy of the board
			var close = CloseMill(location, b, color)
			b.SetPosition(location, color) // Place the color token at the specified location
			if close {
				GenerateRemove(b, &L) // Generate removal states if placing this piece closes a mill
			} else {
				L = append(L, b) // Otherwise, add the board state to the list L
			}
		}
	}

	return L
}

// StaticEstimateOpeningNaive computes a static estimate for an opening board state
// based on the difference between the number of white pieces and black pieces.
func StaticEstimateOpeningNaive(board *MorrisBoard) int {
	numWhitePieces, numBlackPieces := 0, 0

	// Count the number of white and black pieces on the board
	for location := 0; location < 21; location++ {
		switch board.GetPosition(location) {
		case White:
			numWhitePieces++
		case Black:
			numBlackPieces++
		}
	}

	// Calculate the static estimate for the opening phase
	return numWhitePieces - numBlackPieces
}