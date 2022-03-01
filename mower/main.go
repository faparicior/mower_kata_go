package main

import (
	"example.kata.local/mower/internal/mowers/ui"
	"flag"
	"os"
)

func main() {
	flag.Parse()
	os.Exit(ui.Cmd(os.Stdout, flag.Args()))
}
