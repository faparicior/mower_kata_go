package movemowers

import (
	"example.kata.local/mower/mowers/domain/entities"
	"example.kata.local/mower/mowers/domain/valueobjects"
	"example.kata.local/mower/mowers/infrastructure"
)

func MoveMowersHandle(command MoveMowersCommand) MoveMowersResponse {

	instructions := infrastructure.BuildFlatFileInstructionsProvider()
	instructions, _ = instructions.Load(command.filename())

	result := ""

	for i := 0; i < instructions.CountMowers(); i++ {
		mowerId, _ := valueobjects.BuildMowerId("123")
		mowerPosition, _ := instructions.MowerPosition(i)

		mower, _ := entities.BuildMower(mowerId, mowerPosition)
		movements := instructions.MowerMovements(i)
		surface, _ := instructions.Surface()

		for _, movement := range movements {
			uniqueMovement, _ := valueobjects.BuildMowerMovement(movement)
			_ = mower.Move(uniqueMovement, surface)
		}

		result += mower.PositionAsString() + "\n"
	}

	return BuildMoveMowersResponse(result)
}
