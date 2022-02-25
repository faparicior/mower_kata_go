package main

import (
	"example.kata.local/mower/mowers/domain/valueobjects"
	"fmt"
)

func main() {
	var position, _ = valueobjects.BuildXPosition(15)

	fmt.Print(position.Value())
}
