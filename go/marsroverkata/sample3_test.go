package marsrover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFollowInstructions(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{1, 2}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}

	commands := []Command{B, F, L, F, F, R}
	marsRover.acceptCommands(commands)

	assert.Equal(t, "5 2 N", marsRover.currentLocation())
}
