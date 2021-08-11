package marsrover

import (
	"strconv"
)

type Coordinates struct {
	x int
	y int
}

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

func (d Direction) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
}

type Command int

const (
	B Command = iota
	F
	L
	R
)

func (c Command) String() string {
	return [...]string{"B", "F", "L", "R"}[c]
}

type Obstacle struct {
	position Coordinates
}

type Plateau struct {
	maxX      int
	maxY      int
	obstacles []Obstacle
}

type Status int

const (
	OK Status = iota
	NOK
)

func (s Status) String() string {
	return [...]string{"OK", "NOK"}[s]
}

type MarsRover struct {
	plateau  Plateau
	heading  Direction
	position Coordinates
	status   Status
}

func (r *MarsRover) currentLocation() interface{} {
	location := strconv.Itoa(r.position.x) + " " + strconv.Itoa(r.position.y) + " " + r.heading.String()
	return location
}

func (r *MarsRover) acceptCommands(commands []Command) string {
	output := r.PrintRover()
	for _, command := range commands {
		switch command {
		case 0:
			r.backward()
		case 1:
			r.forward()

		case 2:
			r.turnLeft()
		case 3:
			r.turnRight()
		}
		output += r.PrintRover()
	}
	return output
}

func (r *MarsRover) coordinates() Coordinates {
	return r.position
}

func (r *MarsRover) forward() {
	newCoordinates := r.position
	switch r.heading {
	case 0:
		if r.position.y == r.plateau.maxY {
			newCoordinates.y = 0
		} else {
			newCoordinates.y += 1
		}
	case 1:
		if r.position.x == r.plateau.maxX {
			newCoordinates.x = 0
		} else {
			newCoordinates.x += 1
		}
	case 2:
		if r.position.y == 0 {
			newCoordinates.y = r.plateau.maxY
		} else {
			newCoordinates.y -= 1
		}
	case 3:
		if r.position.x == 0 {
			newCoordinates.x = r.plateau.maxX
		} else {
			newCoordinates.x -= 1
		}
	}
	r.checkForObstacles(newCoordinates)
}

func (r *MarsRover) backward() {
	newCoordinates := r.position
	switch r.heading {
	case 0:
		if r.position.y == 0 {
			newCoordinates.y = r.plateau.maxY
		} else {
			newCoordinates.y -= 1
		}
	case 1:
		if r.position.x == 0 {
			newCoordinates.x = r.plateau.maxX
		} else {
			newCoordinates.x -= 1
		}
	case 2:
		if r.position.y == r.plateau.maxY {
			newCoordinates.y = 0
		} else {
			newCoordinates.y += 1
		}
	case 3:
		if r.position.x == r.plateau.maxX {
			newCoordinates.x = 0
		} else {
			newCoordinates.x += 1
		}
	}
	r.checkForObstacles(newCoordinates)
}

func (r *MarsRover) checkForObstacles(coordinates Coordinates) {
	for _, obstacle := range r.plateau.obstacles {
		if obstacle.position.x == coordinates.x && obstacle.position.y == coordinates.y {
			r.status = 1
			coordinates = r.position
		}
	}
	r.position = coordinates
}

func (r *MarsRover) turnRight() {
	if r.heading < 3 {
		r.heading += 1
	} else {
		r.heading = 0
	}

}

func (r *MarsRover) turnLeft() {
	if r.heading > 0 {
		r.heading -= 1
	} else {
		r.heading = 3
	}

}
