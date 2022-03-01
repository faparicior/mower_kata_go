package infrastructure

import (
	"bufio"
	"example.kata.local/mower/mowers/domain/valueobjects"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FlatFileInstructionsProvider struct {
	filename  string
	surface   string
	positions []string
	movements []string
}

func BuildFlatFileInstructionsProvider() FlatFileInstructionsProvider {
	return FlatFileInstructionsProvider{}
}

func (provider FlatFileInstructionsProvider) Load(filename string) (FlatFileInstructionsProvider, error) {
	var surface string

	csvFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(fmt.Errorf("error opening file: %v", err))
		return FlatFileInstructionsProvider{}, err
	}

	reader := bufio.NewScanner(csvFile)

	surface, err = readSurface(reader)
	if err != nil {
		return FlatFileInstructionsProvider{}, err
	}

	positions, movements, instructionsProvider, err := readPositionsAndMovements(reader)

	if err != nil {
		return instructionsProvider, err
	}

	return FlatFileInstructionsProvider{
		filename:  filename,
		surface:   surface,
		positions: positions,
		movements: movements,
	}, nil
}

func (provider FlatFileInstructionsProvider) Surface() (valueobjects.Surface, error) {
	surfaceAsArray := strings.Split(provider.surface, " ")
	value, _ := strconv.Atoi(surfaceAsArray[0])
	xSize, _ := valueobjects.BuildSurfaceXSize(value)
	value, _ = strconv.Atoi(surfaceAsArray[1])
	ySize, _ := valueobjects.BuildSurfaceYSize(value)

	surface, _ := valueobjects.BuildSurface(xSize, ySize)

	return surface, nil
}

func (provider *FlatFileInstructionsProvider) MowerPosition(index int) (valueobjects.MowerPosition, error) {
	const xPositionIndex int = 0
	const yPositionIndex int = 1
	const orientationIndex int = 2

	position := strings.Split(provider.positions[index], " ")

	value, _ := strconv.Atoi(position[xPositionIndex])
	xPosition, _ := valueobjects.BuildXPosition(value)
	value, _ = strconv.Atoi(position[yPositionIndex])
	yPosition, _ := valueobjects.BuildYPosition(value)
	orientation, _ := valueobjects.BuildMowerOrientation(position[orientationIndex])

	return valueobjects.BuildMowerPosition(xPosition, yPosition, orientation)
}

func (provider *FlatFileInstructionsProvider) MowerMovements(index int) []valueobjects.MowerMovement {
	movementsAsArray := strings.Split(provider.movements[index], "")

	var mowerMovements []valueobjects.MowerMovement

	for _, movement := range movementsAsArray {
		mowerMovement, _ := valueobjects.BuildMowerMovement(movement)
		mowerMovements = append(mowerMovements, mowerMovement)
	}

	return mowerMovements
}

func (provider *FlatFileInstructionsProvider) CountMowers() int {
	return len(provider.positions)
}

func readPositionsAndMovements(reader *bufio.Scanner) ([]string, []string, FlatFileInstructionsProvider, error) {
	var positions []string
	var movements []string

	for reader.Scan() {
		positions = append(positions, reader.Text())
		reader.Scan()
		movements = append(movements, reader.Text())
	}
	if err := reader.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input for Instructions:", err)
		return nil, nil, FlatFileInstructionsProvider{}, err
	}
	return positions, movements, FlatFileInstructionsProvider{}, nil
}

func readSurface(reader *bufio.Scanner) (string, error) {
	reader.Scan()
	if err := reader.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input for Surface instructions:", err)
		return "", err
	}

	return reader.Text(), nil
}
