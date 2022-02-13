package main

import (
	"fmt"

	"example.kata.local/mower/mowers/valueobject"
)

func main() {
	var position, _ = XPosition.Build(15)

	fmt.Print(position.Value())
}
