package marsrover

import (
	"testing"

	approvaltests "github.com/approvals/go-approval-tests"
)

func TestPrintPlateau(t *testing.T) {
	rocks := []Obstacle{
		{position: Coordinates{5, 3}},
		{position: Coordinates{1, 1}},
		{position: Coordinates{3, 4}},
	}
	plateau := Plateau{maxX: 5, maxY: 5, obstacles: rocks}

	printer := generateGrid(plateau)
	printer.addObstacles(plateau.obstacles)

	approvaltests.VerifyString(t, plateau.PrintPlateau(printer))

}
