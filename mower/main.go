package main

import (
	"fmt"

	"example.kata.local/mower/mowers/valueobject"
)

func main() {
	var culo XPosition.Instance = XPosition.Build(15)

	fmt.Print(culo.Value())
}
