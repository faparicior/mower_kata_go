package infrastructure

import (
	"bufio"
	"example.kata.local/mower/mowers/domain/valueobjects"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BuildFlatFileInstructionsProvider() Instructions {
	return Instructions{
		surface:      "",
		instructions: nil,
	}
}

func (instructions Instructions) Load(filename string) (Instructions, error) {
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(fmt.Errorf("error opening file: %v", err))
		return Instructions{}, err
	}

	reader := bufio.NewScanner(csvFile)
	err = readSurface(reader, &instructions)
	if err != nil {
		return Instructions{}, err
	}

	for reader.Scan() {
		instructions.instructions = append(instructions.instructions, reader.Text())
	}
	if err := reader.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input for Instructions:", err)
		return Instructions{}, err
	}

	return instructions, nil
}

func readSurface(reader *bufio.Scanner, instructions *Instructions) error {
	reader.Scan()
	if err := reader.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input for Surface instructions:", err)
		return err
	}

	instructions.surface = reader.Text()

	return nil
}

func (instructions Instructions) Surface() (valueobjects.Surface, error) {
	surface := strings.Split(instructions.surface, " ")

	//TODO: Validation to assure two values
	value, _ := strconv.Atoi(surface[0])
	xSurface, _ := valueobjects.BuildSurfaceXSize(value)
	value, _ = strconv.Atoi(surface[1])
	ySurface, _ := valueobjects.BuildSurfaceYSize(value)

	return valueobjects.BuildSurface(xSurface, ySurface)
}

func (instructions Instructions) MowerPosition(index int) (valueobjects.MowerPosition, error) {
	const orientationIndex int = 2
	const xPositionIndex int = 0
	const yPositionIndex int = 1

	position := strings.Split(instructions.instructions[realIndex(index)], " ")

	value, _ := strconv.Atoi(position[xPositionIndex])
	xPosition, _ := valueobjects.BuildXPosition(value)
	value, _ = strconv.Atoi(position[yPositionIndex])
	yPosition, _ := valueobjects.BuildYPosition(value)
	orientation, _ := valueobjects.BuildMowerOrientation(position[orientationIndex])

	return valueobjects.BuildMowerPosition(xPosition, yPosition, orientation)
}

func (instructions Instructions) MowerMovements(index int) []string {
	return strings.Split(instructions.instructions[realIndex(index)+1], "")
}

func realIndex(index int) int {
	const indexCompensation int = 1
	const firstElement int = 0

	if firstElement == index {
		return 0
	}

	return index + indexCompensation
}
