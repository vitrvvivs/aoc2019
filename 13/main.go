package main

import (
	"../geometry"
	"../intcode"
	"../utils"
	"fmt"
	"os"
	"os/exec"
)

const EMPTY = 0
const WALL = 1
const BLOCK = 2
const PADDLE = 3
const BALL = 4

var mapping = map[int]rune{EMPTY: ' ', WALL: '█', BLOCK: '■', PADDLE: '─', BALL: 'o'}

const INTERACTIVE = false
const LEFT = -1
const RIGHT = 1
const NEUTRAL = 0
const RIGHTARROW = 67
const LEFTARROW = 68

type game struct {
	geometry.Point
	intcode.Program
	data       [][]int
	outputMode int
	score      int
	ball       geometry.Point
	paddle     geometry.Point
}

func newGame() *game {
	w, h := 38, 21
	g := game{}
	g.Program = *intcode.NewProgram("13/input.txt")
	g.Program.Output = g.input
	if INTERACTIVE {
		g.Program.Input = g.interactiveOutput
	} else {
		g.Program.Input = g.output
	}
	g.outputMode = 0
	g.data = make([][]int, w)
	for i := range g.data {
		g.data[i] = make([]int, h)
	}
	return &g
}

func (g game) print() {
	fmt.Print("\033[2J")
	geometry.PrintMapping(g.data, mapping)
}
func (g *game) input(val int) {
	switch g.outputMode {
	case 0:
		g.X = val
	case 1:
		g.Y = val
	case 2:
		if g.X == -1 && g.Y == 0 {
			g.score = val
		} else {
			g.data[g.X][g.Y] = val
			if val == BALL {
				g.ball = g.Point
			}
			if val == PADDLE {
				g.paddle = g.Point
			}
		}
	}
	g.outputMode += 1
	g.outputMode %= 3
}
func (g *game) output() int { // to Program
	return utils.Limit(g.ball.X-g.paddle.X, 1)
}

func (g *game) interactiveOutput() int { // to Program and screen
	b := make([]byte, 3)
	g.print()
	os.Stdin.Read(b)
	fmt.Println(b)
	switch b[2] {
	case LEFTARROW:
		return LEFT
	case RIGHTARROW:
		return RIGHT
	default:
		return NEUTRAL
	}
}

func main() {
	// PART A
	game := newGame()
	game.Run()

	count := 0
	for _, line := range game.data {
		for _, tile := range line {
			if tile == BLOCK {
				count++
			}
		}
	}
	fmt.Println(count)

	// PART B
	game = newGame()
	game.State[0] = 2
	if INTERACTIVE {
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	}
	game.Run()
	fmt.Println(game.score)
}
