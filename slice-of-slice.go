package main

import (
	"fmt"
	"strings"
)

func main() {
	// Tic Tac Toe
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[1][1] = "X"
	board[2][2] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Printf("%s\n", strings.Join(board[2], "	"))
}
