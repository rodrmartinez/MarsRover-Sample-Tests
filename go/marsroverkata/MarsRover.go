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

func (r *MarsRover) acceptCommands(commands []Command) {
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
	}

}

func (r *MarsRover) coordinates() Coordinates {
	return r.position
}

func (r *MarsRover) forward() {
	switch r.heading {
	case 0:
		r.position.y += 1
	case 1:
		r.position.x += 1
	case 2:
		r.position.y -= 1
	case 3:
		r.position.x -= 1
	}
}

func (r *MarsRover) backward() {
	switch r.heading {
	case 0:
		r.position.y -= 1
	case 1:
		r.position.x -= 1
	case 2:
		r.position.y += 1
	case 3:
		r.position.x += 1
	}

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
