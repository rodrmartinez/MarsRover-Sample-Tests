package marsrover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFollowAllInstructions(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{1, 2}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}

	commands := []Command{B, F, L, F, F, R}
	for _, command := range commands {
		switch command {
		case B:
			marsRover.backward()
		case F:
			marsRover.forward()
		case L:
			marsRover.turnLeft()
		case R:
			marsRover.turnRight()
		}
	}
	expectedPosition := Coordinates{5, 2}
	assert.Equal(t, expectedPosition, marsRover.coordinates())
	assert.Equal(t, N, marsRover.heading)
}
