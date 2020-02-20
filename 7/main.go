package main

import (
	"../intcode"
	"fmt"
)

const DEBUG = false

type T interface{}

func combinations(elements []T) chan []T {
	ch := make(chan []T)
	go func() {
		_combinations(make([]T, 0), elements, ch)
		close(ch)
	}()
	return ch
}

func _combinations(used []T, available []T, ch chan []T) {
	if len(available) == 1 {
		ch <- append(used, available[0])
		return
	}
	for i, v := range available {
		next := make([]T, len(available))
		copy(next, available)
		_combinations(append(used, v), append(next[:i], next[i+1:]...), ch)
	}
}

func max_output(statefile string, elements []int) int {
	_elements := make([]T, len(elements))
	for i, v := range elements {
		_elements[i] = v
	}

	max := 0
	for seq := range combinations(_elements[:]) {
		output := 0
		for _, phase := range seq {
			inputs := []int{phase.(int), output}
			program := intcode.NewProgram(statefile)
			program.Input = func() int { r := inputs[0]; inputs = inputs[1:]; return r }
			program.Output = func(x int) { output = x }
			program.Run()
		}
		if output > max {
			max = output
		}
	}
	return max
}

func feedback_loop(statefile string, elements []int) int {
	done := make(chan bool)
	_elements := make([]T, len(elements))
	for i, v := range elements {
		_elements[i] = v
	}

	max := 0
	for seq := range combinations(_elements[:]) {
		output := 0
		// initialize programs
		channels := make([]chan int, 5)
		programs := make([]*intcode.Program, 5)
		for i := 0; i < 5; i++ {
			programs[i], channels[i] = intcode.NewProgram(statefile), make(chan int)
		}
		// connect outputs and inputs
		for i, _ := range seq {
			i := i
			programs[i].Input = func() int { x := <-channels[(i+4)%5]; return x }
			programs[i].Output = func(x int) { channels[i] <- x }
			programs[i].Halt = func() { close(channels[i]) }
		}
		// second-to-last program signals done, so the last output of the last program is gotten by the main thread
		programs[3].Halt = func() { close(channels[3]); done <- true }
		for i, v := range seq {
			i, v := i, v.(int)
			go programs[i].Run()
			channels[(i+4)%5] <- v
		}

		channels[4] <- 0 // read by program[0]
		<-done
		output = <-channels[4]

		if output > max {
			max = output
		}
	}
	close(done)
	return max
}

func main() {
	if DEBUG {
		fmt.Println(max_output("7/test1.txt", []int{0, 1, 2, 3, 4}), 43210)
		fmt.Println(max_output("7/test2.txt", []int{0, 1, 2, 3, 4}), 54321)
		fmt.Println(max_output("7/test3.txt", []int{0, 1, 2, 3, 4}), 65210)
		fmt.Println(feedback_loop("7/test4.txt", []int{5, 6, 7, 8, 9}), 139629729)
		fmt.Println(feedback_loop("7/test5.txt", []int{5, 6, 7, 8, 9}), 18216)
	} else {
		fmt.Println(max_output("7/input.txt", []int{0, 1, 2, 3, 4}))
		fmt.Println(feedback_loop("7/input.txt", []int{5, 6, 7, 8, 9}))
	}
}
