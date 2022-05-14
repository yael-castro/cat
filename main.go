package main

import (
	"fmt"

	"github.com/yael-castro/cat/internal/game"
)

func main() {
	g := game.New()

	var (
		state    = game.NoState
		position = game.Position(0)
	)

	fmt.Println(g)

	for !state.Is(game.Draw | game.Player1Won | game.Player2Won) {
		if g.Turn() {
			fmt.Print("\nPlayer 1: ")
		} else {
			fmt.Print("\nPlayer 2: ")
		}

		fmt.Scan(&position)

		state = g.Play(1 << position >> 1)
		if state.Is(game.InvalidTurn) {
			fmt.Printf(`Position "%v" does not exists`, position)
		}

		if state.Is(game.NoSpace) {
			fmt.Printf(`Position "%v" is already taken`, position)
		}

		fmt.Println(g)
	}

	switch state {
	case game.Player1Won:
		fmt.Println("\nPlayer 1 won!")
	case game.Player2Won:
		fmt.Println("\nPlayer 2 won!")
	default:
		fmt.Println("\nDraw!")
	}
}
