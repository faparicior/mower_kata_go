package ui

import (
	movemowers2 "example.kata.local/mower/internal/mowers/application/movemowers"
	"fmt"
	"io"
)

func Cmd(out io.Writer, args []string) int {
	if len(args) == 0 {
		fmt.Fprintf(out, "Missing filename\n")
		return 1
	}

	moveMowersResponse := movemowers2.MoveMowersHandle(movemowers2.BuildMoveMowersCommand(args[0]))

	fmt.Fprintf(out, moveMowersResponse.Response())
	return 0
}
