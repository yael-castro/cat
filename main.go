package main

import (
	"fmt"

	"github.com/yael-castro/tic-tac-toe/internal/game"
)

func main() {
	g := game.New()

	var (
		state         = game.Continue
		position uint = 0
	)

	fmt.Println(g)

	for state.Is(game.Continue) {
		if g.Turn() {
			fmt.Print("\nPlayer 1: ")
		} else {
			fmt.Print("\nPlayer 2: ")
		}

		_, err := fmt.Scan(&position)
		if err != nil {
			fmt.Printf("%v\n%v\n", err, position)
			continue
		}

		state = g.Play(position)
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
	case game.Player1Won | game.Player2Won:
		fmt.Println("\nDraw!")
	}
}
