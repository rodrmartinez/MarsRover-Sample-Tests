package marsrover

import (
	"testing"

	approvaltests "github.com/approvals/go-approval-tests"
	"github.com/stretchr/testify/assert"
)

func TestObstacleForward(t *testing.T) {
	rocks := []Obstacle{
		{position: Coordinates{0, 3}},
	}
	plateau := Plateau{maxX: 5, maxY: 5, obstacles: rocks}
	startingPosition := Coordinates{0, 0}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}

	commands := []Command{F, F, F, F, F, F}

	approvaltests.VerifyString(t, marsRover.acceptCommands(commands))
	assert.Equal(t, "NOK", marsRover.status.String())
	assert.Equal(t, "0 2 N", marsRover.currentLocation())
}

func TestObstacleBackward(t *testing.T) {
	rocks := []Obstacle{
		{position: Coordinates{5, 3}},
	}
	plateau := Plateau{maxX: 5, maxY: 5, obstacles: rocks}
	startingPosition := Coordinates{5, 5}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}

	commands := []Command{B, B, B, B, B, B}

	approvaltests.VerifyString(t, marsRover.acceptCommands(commands))
	assert.Equal(t, "NOK", marsRover.status.String())
	assert.Equal(t, "5 4 N", marsRover.currentLocation())
}

func TestObstacleBackward2(t *testing.T) {
	rocks := []Obstacle{
		{position: Coordinates{2, 5}},
	}
	plateau := Plateau{maxX: 5, maxY: 5, obstacles: rocks}
	startingPosition := Coordinates{5, 5}
	marsRover := MarsRover{plateau: plateau, heading: E, position: startingPosition}

	commands := []Command{B, B, B, B, B, B}

	approvaltests.VerifyString(t, marsRover.acceptCommands(commands))
	assert.Equal(t, "NOK", marsRover.status.String())
	assert.Equal(t, "3 5 E", marsRover.currentLocation())
}
