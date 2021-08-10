package marsrover

import (
	"testing"

	approvaltests "github.com/approvals/go-approval-tests"
	"github.com/stretchr/testify/assert"
)

func TestObstacles(t *testing.T) {
	rocks := []Obstacle{
		{position: Coordinates{6, 6}},
	}
	plateau := Plateau{maxX: 10, maxY: 10, obstacles: rocks}
	startingPosition := Coordinates{0, 0}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}

	commands := []Command{F, F, F, R, F, F, F, R, B, B, B, R, R, R, F, F, F, F}

	approvaltests.VerifyString(t, marsRover.acceptCommands(commands))
	assert.Equal(t, "NOK", marsRover.status.String())
	assert.Equal(t, "0 2 N", marsRover.currentLocation())
}
