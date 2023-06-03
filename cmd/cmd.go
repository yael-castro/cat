package main

import (
	"fmt"
	"github.com/yael-castro/xogo"
)

func main() {
	var (
		position uint
		g        xogo.Game
		state    = xogo.Continue
	)

	fmt.Println(g)

	for state.Is(xogo.Continue) {
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
		if state.Is(xogo.InvalidTurn) {
			fmt.Printf(`Position "%v" does not exists`, position)
		}

		if state.Is(xogo.NoSpace) {
			fmt.Printf(`Position "%v" is already taken`, position)
		}

		fmt.Println(g)
	}

	switch state {
	case xogo.Player1Won:
		fmt.Println("\nPlayer 1 won!")
	case xogo.Player2Won:
		fmt.Println("\nPlayer 2 won!")
	case xogo.Player1Won | xogo.Player2Won:
		fmt.Println("\nDraw!")
	}
}
