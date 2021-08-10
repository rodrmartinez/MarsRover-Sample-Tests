package marsrover

import (
	"testing"

	approvaltests "github.com/approvals/go-approval-tests"
	"github.com/stretchr/testify/assert"
)

func TestPrintRover(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{1, 2}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}
	commands := []Command{F, B, L, F, F, F}

	approvaltests.VerifyString(t, marsRover.acceptCommands(commands))
	assert.Equal(t, "4 2 W", marsRover.currentLocation())
}
