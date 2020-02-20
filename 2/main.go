package main

import (
	"../intcode"
	"fmt"
)

const LIMIT = 100

func a() int {
	program := intcode.NewProgram("2/input.txt")
	program.State[1] = 12
	program.State[2] = 2
	program.Run()
	return program.State[0]
}
func b() int {
	for i := 0; i < LIMIT; i++ {
		for j := 0; j < LIMIT; j++ {
			program := intcode.NewProgram("2/input.txt")
			program.State[1] = i
			program.State[2] = j
			program.Run()
			if program.State[0] == 19690720 {
				return i*100 + j
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(a())
	fmt.Println(b())
}
