package main

import (
	"../intcode"
)

const DEBUG = false

func main() {
	if DEBUG {
		intcode.NewProgram("9/test1.txt").Run()
		intcode.NewProgram("9/test2.txt").Run()
		intcode.NewProgram("9/test3.txt").Run()
	} else {
		program := intcode.NewProgram("9/input.txt")
		program.Input = func() int { return 1 }
		program.Run()

		program = intcode.NewProgram("9/input.txt")
		program.Input = func() int { return 2 }
		program.Run()
	}
}
