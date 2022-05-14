// Package game contains the rules and everything needed to play the Cat's Game
package game

import "fmt"

// Position indicates a position on Board
type Position uint

const (
	// First row
	// A indicates the first row and the first column
	A Position = 1 << iota
	// B indicates the first row and the second column
	B
	// C indicates the first row and the third column
	C
	// Second row
	// D indicates the second row and the first column
	D
	// E indicates the second row and the second column
	E
	// F indicates the second row and the third column
	F
	// Third row
	// G indicates the third row and the first column
	G
	// H indicates the third row and the second column
	H
	// I indicates the third row and the third column
	I
)

// Board defines the bit mask for a Cat's Game player
type Board uint

// IsFull returns true if all positions on the board all full
func (b Board) IsFull() bool {
	return int(b) == int(A|B|C|D|E|F|G|H|I)
}

// contains indicates if a space on the board is occuped
func (b Board) contains(p Position) bool {
	return int(b)&int(p) == int(p)
}

// IsComplete indicates if the board has been completed
func (b Board) IsComplete() bool {
	switch {
	// Horizontal cases
	case b.contains(A | B | C):
		return true

	case b.contains(D | E | F):
		return true

	case b.contains(G | H | I):
		return true

	// Vertical cases
	case b.contains(A | D | G):
		return true

	case b.contains(B | E | H):
		return true

	case b.contains(C | F | I):
		return true

	// Slash cases
	case b.contains(A | E | I):
		return true

	case b.contains(C | E | G):
		return true
	}

	return false
}

// State indicates the game state
type State int

// Is indicates if the current State match to composite state
func (s State) Is(state State) bool {
	return s&state == s
}

// Supported values for State, defines the
const (
	// NoState must be used in the game begin
	NoState State = 1 << iota
	// Draw indicates the
	Draw
	// Continue indicates the success turn
	Continue
	// Player1Won indicates the player 1 victory
	Player1Won
	// Player2Won indicates the player 2 victory
	Player2Won
	// InvalidTurn indicates the out of range position
	InvalidTurn
	// NoSpace indicates that some space is already occupied
	NoSpace
)

// Game controls the game flow and the game rules
type Game struct {
	turn    bool
	player1 Board
	player2 Board
}

// New returns an instance of Game
func New() *Game {
	return &Game{}
}

// Turn returns the current player turn
//
// If returns true, is the turn of player 1
//
// If returns false, is the turn of player 2
func (g Game) Turn() bool {
	return !g.turn
}

// Player1 returns the Board of player 1
func (g Game) Player1() Board {
	return g.player1
}

// Player2 returns the Board of player 2
func (g Game) Player2() Board {
	return g.player2
}

// validatePosition check if the position is valid
func (g Game) validatePosition(p Position) bool {
	switch p {
	case A, B, C, D, E, F, G, H, I:
		return true
	}

	return false
}

// Play starts the game and continue using the current state
func (g *Game) Play(p Position) State {
	if !g.validatePosition(p) {
		return InvalidTurn
	}

	if (g.player1 | g.player2).contains(p) {
		return NoSpace
	}

	// Player 1 turn
	if g.Turn() {
		g.player1 = g.player1 | Board(p)
		goto checkResults
	}

	// Player 2 turn
	g.player2 = g.player2 | Board(p)

checkResults:
	g.turn = !g.turn

	if g.player1.IsComplete() {
		return Player1Won
	}

	if g.player2.IsComplete() {
		return Player2Won
	}

	if (g.player1 | g.player2).IsFull() {
		return Draw
	}

	return Continue
}

// Reset clean the player boards but maintains the last turn
func (g *Game) Reset() {
	g.player1 = 0
	g.player2 = 0
}

// String returns the string representation for the both boards
func (g Game) String() string {
	out := ""

	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			out += "\n"
		}

		if g.player1.contains(1 << i) {
			out += "  x  "
		} else if g.player2.contains(1 << i) {
			out += "  o  "
		} else {
			out += fmt.Sprintf("  %d  ", i+1)
		}
	}

	return out
}
