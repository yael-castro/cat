// Package game contains the rules and everything needed to play the Cat's Game
package game

import "fmt"

// Board is a bitmask used to represents a player board
//
// The board can be imagined as shown below... Where each space in the board can be selected by position
//
//  0  |  1  |  2
//  3  |  4  |  5
//  6  |  7  |  8
//
// But it is actually a bitmask of 9 bits that use each bit to represents a space in the board
//
// Board position | 8 7 6 5 4 3 2 1 0
// Bitmask        | 0 0 0 0 0 0 0 0 1
//
type Board uint

// IsFull returns true if all positions on the board all full
func (b Board) IsFull() bool {
	return int(b) == int(0b111_111_111)
}

// Contains indicates if a position on the board is already taken
func (b Board) Contains(bit uint) bool {
	b2 := Board(1 << bit)
	return b&b2 == b2
}

// contains indicates if a raw space on the board is occuped
func (b Board) contains(b2 Board) bool {
	return b&b2 == b2
}

// IsComplete indicates if the board has been completed
func (b Board) IsComplete() bool {
	switch {
	// Horizontal cases
	case b.contains(0b000_000_111):
		return true

	case b.contains(0b000_111_000):
		return true

	case b.contains(0b111_000_000):
		return true

	// Vertical cases
	case b.contains(0b001_001_001):
		return true

	case b.contains(0b010_010_010):
		return true

	case b.contains(0b100_100_100):
		return true

	// Slash cases
	case b.contains(0b100_010_001):
		return true

	case b.contains(0b001_010_100):
		return true
	}

	return false
}

// State indicates the game state
type State int

// Is indicates if the current State match to composite state
func (s State) Is(state State) bool {
	return s&state == state
}

// Supported values for State, defines the
const (
	// Continue indicates the success turn
	Continue State = 1 << iota
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

// Play starts the game and continue using the current state
func (g *Game) Play(position uint) State {
	if position > 8 {
		return Continue | InvalidTurn
	}

	b := Board(1 << position)

	if (g.player1 | g.player2).contains(b) {
		return Continue | NoSpace
	}

	// Player 1 turn
	if g.Turn() {
		g.player1 = g.player1 | b
		goto checkResults
	}

	// Player 2 turn
	g.player2 = g.player2 | b

checkResults:
	g.turn = !g.turn

	if g.player1.IsComplete() {
		return Player1Won
	}

	if g.player2.IsComplete() {
		return Player2Won
	}

	if (g.player1 | g.player2).IsFull() {
		return Player1Won | Player2Won
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

	for i := uint(0); i < 9; i++ {
		if i%3 == 0 {
			out += "\n"
		}

		if g.player1.Contains(i) {
			out += "  x  "
		} else if g.player2.Contains(i) {
			out += "  o  "
		} else {
			out += fmt.Sprintf("  %d  ", i)
		}
	}

	return out
}
