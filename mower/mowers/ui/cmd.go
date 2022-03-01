package ui

import (
	"example.kata.local/mower/mowers/application/movemowers"
	"fmt"
	"io"
)

func Cmd(out io.Writer, args []string) int {
	if len(args) == 0 {
		fmt.Fprintf(out, "Missing filename\n")
		return 1
	}

	moveMowersResponse := movemowers.MoveMowersHandle(movemowers.BuildMoveMowersCommand(args[0]))

	fmt.Fprintf(out, moveMowersResponse.Response())
	return 0
}
