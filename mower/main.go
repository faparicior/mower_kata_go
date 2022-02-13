package main

import (
	"fmt"

	"example.kata.local/mower/mowers/valueobject"
)

func main() {
	var position, _ = valueobject.BuildXPosition(15)

	fmt.Print(position.Value())
}
