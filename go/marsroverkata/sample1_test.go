package marsrover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanRotateLeft(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{1, 2}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}

	marsRover.turnLeft()

	assert.Equal(t, W, marsRover.heading)

	marsRover.turnLeft()

	assert.Equal(t, S, marsRover.heading)

	marsRover.turnLeft()

	assert.Equal(t, E, marsRover.heading)
	marsRover.turnLeft()

	assert.Equal(t, N, marsRover.heading)
}

func TestCanRotateRight(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{1, 2}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}

	marsRover.turnRight()
	assert.Equal(t, E, marsRover.heading)

	marsRover.turnRight()

	assert.Equal(t, S, marsRover.heading)

	marsRover.turnRight()

	assert.Equal(t, W, marsRover.heading)
	marsRover.turnRight()

	assert.Equal(t, N, marsRover.heading)
}
