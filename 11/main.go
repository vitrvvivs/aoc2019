package main

import (
	"../geometry"
	"../intcode"
	"fmt"
)

const DEBUG = false

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
func itob(i int) bool {
	return i > 0
}

// ROBOT - moving and painting
const PAINT = 0
const ROTATE = 1

var DIRECTIONS = [4]geometry.Vector{
	{0, -1, 0}, // UP
	{1, 0, 0},  // RIGHT
	{0, 1, 0},  // DOWN
	{-1, 0, 0}, // LEFT
}
var DIRECTIONDEBUG = [4]string{"^", ">", "V", "<"}

type robot struct {
	geometry.Point
	direction int
}

func (r robot) getDirection() geometry.Vector {
	return DIRECTIONS[r.direction]
}
func (r *robot) rotate(dir int) {
	// direction is either 0 (LEFT) or 1 (RIGHT)
	r.direction += dir
	r.direction %= 4
	if r.direction < 0 {
		r.direction += 4
	}
	// alway move after rotating
	r.X += r.getDirection().X
	r.Y += r.getDirection().Y
}

// SHIP - surface painted
type ship struct {
	geometry.Box
	color   [][]bool
	painted [][]bool
}

func newShip(b geometry.Box) ship {
	s := ship{}
	s.Box = b
	s.color = make([][]bool, b.W)
	s.painted = make([][]bool, b.W)
	for i := range s.color {
		s.color[i] = make([]bool, b.H)
		s.painted[i] = make([]bool, b.H)
	}
	return s
}
func (s ship) Print(r robot) {
	for y := range s.color[0] {
		for x := range s.color {
			if x == r.X && y == r.Y {
				fmt.Print(DIRECTIONDEBUG[r.direction])
			} else if s.color[x][y] {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	r := robot{geometry.Point{X: 100, Y: 100}, 0}
	s := newShip(geometry.Box{W: 200, H: 200})
	outputMode := PAINT

	program := intcode.NewProgram("11/input.txt")
	program.Input = func() int {
		return btoi(s.color[r.X][r.Y])
	}
	program.Output = func(i int) {
		if DEBUG {
			fmt.Print(i)
		}
		if outputMode == PAINT {
			s.color[r.X][r.Y] = itob(i)
			s.painted[r.X][r.Y] = true
			outputMode = ROTATE
		} else if outputMode == ROTATE {
			if i == 0 {
				i = -1
			}
			r.rotate(i)
			outputMode = PAINT
			if DEBUG {
				s.Print(r)
			}
		}
	}
	program.Halt = func() {
		count := 0
		for _, line := range s.painted {
			for _, painted := range line {
				if painted {
					count++
				}
			}
		}
		fmt.Println(count)
	}

	program.Run()

	r = robot{geometry.Point{X: 0, Y: 0}, 0}
	s = newShip(geometry.Box{W: 43, H: 6})
	s.color[r.X][r.Y] = true
	program.Reload()
	program.Halt = func() {
		s.Print(r)
	}
	program.Run()
}
