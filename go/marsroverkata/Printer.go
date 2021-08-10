package marsrover

import (
	"strconv"
)

type Grid [][]string

// Printer ..
type Printer interface {
	PrintPlateau() string
	PrintRover() string
}

func (p Plateau) PrintPlateau(grid Grid) string {
	axisX := generateAxis(p.maxX)
	axisY := generateAxis(p.maxY)
	result := ""

	for i := len(axisY) - 1; i >= 0; i-- {
		result += axisY[i] + "\t"
		for j := range axisY {
			result += grid[j][i] + "\t"
		}
		result += "\n"
	}
	result += "\t"
	for x := range axisX {
		result += axisX[x] + "\t"
	}
	return result
}

func (g Grid) addObstacles(obstacles []Obstacle) {
	for _, obstacle := range obstacles {
		g[obstacle.position.x][obstacle.position.y] = "*"
	}
}

func generateGrid(p Plateau) Grid {
	grid := make(Grid, p.maxX+1)
	for i := range grid {
		grid[i] = make([]string, p.maxY+1)
	}
	return grid
}

func generateAxis(max int) []string {
	axis := make([]string, max+1)
	for i := range axis {
		axis[i] = strconv.Itoa(i)
	}
	return axis
}
