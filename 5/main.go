package main

import (
	"../intcode"
)

func main() {
	program := intcode.NewProgram("5/input.txt")
	program.Input = func() int { return 1 }
	program.Run()

	program = intcode.NewProgram("5/input.txt")
	program.Input = func() int { return 5 }
	program.Run()
}
