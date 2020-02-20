package intcode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DEBUG = false

const ADD = 1
const MULTIPLY = 2
const INPUT = 3
const OUTPUT = 4
const JUMP_IF_TRUE = 5
const JUMP_IF_FALSE = 6
const LESS_THAN = 7
const EQUALS = 8
const ADD_RELATIVE = 9
const HALT = 99

var NUM_ARGS = map[int]int{
	ADD:           3,
	MULTIPLY:      3,
	INPUT:         1,
	OUTPUT:        1,
	JUMP_IF_TRUE:  2,
	JUMP_IF_FALSE: 2,
	LESS_THAN:     3,
	EQUALS:        3,
	ADD_RELATIVE:  1,
}

const POSITION = 0
const IMMEDIATE = 1
const RELATIVE = 2

type Program struct {
	hooks
	statefile string
	State     []int
	pc        int
	rb        int
}
type hooks struct {
	Output func(int)
	Input  func() int
	Halt   func()
}

func NewProgram(statefile string) *Program {
	this := &Program{statefile: statefile}

	this.Reload()
	this.Input = func() int { var r int; fmt.Scanf("%d\n", &r); return r }
	this.Output = func(x int) { fmt.Println(x) }
	this.Halt = func() {}
	return this
}
func (p *Program) Reload() {
	file, _ := os.Open(p.statefile)
	fileInfo, _ := file.Stat()
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	// READ DATA
	i := 0
	state := make([]int, fileInfo.Size()-1)
	for _, num := range strings.Split(scanner.Text(), ",") {
		state[i], _ = strconv.Atoi(num)
		i++
	}
	state = state[:i]

	p.State = append(state, make([]int, 2048)...)
}
func (p Program) Run() {
	p.pc = 0
	p.rb = 0
	for p.State[p.pc] != HALT {
		p.runNextInstruction()
	}
	p.Halt()
}

func (p *Program) runNextInstruction() {
	/* ABCDE
	    2105
	DE - two-digit opcode,      05 == opcode 2
	 C - mode of 1st parameter,  1 == immediate mode
	 B - mode of 2nd parameter,  2 == relative mode
	 A - mode of 3rd parameter,  0 == position mode,
	                                  omitted due to being a leading zero
	*/
	op := p.State[p.pc] % 100
	modes := [3]int{p.State[p.pc] / 100 % 10, p.State[p.pc] / 1000 % 10, p.State[p.pc] / 10000 % 10}
	var args, values [3]int
	for i, mode := range modes[:NUM_ARGS[op]] {
		arg := p.State[p.pc+i+1]
		if mode == RELATIVE {
			arg = p.rb + arg
		}
		if mode == POSITION || mode == RELATIVE {
			values[i] = p.State[arg]
		} else if mode == IMMEDIATE {
			values[i] = arg
		}
		args[i] = arg
	}

	switch op {
	case ADD:
		p.State[args[2]] = values[0] + values[1]
	case MULTIPLY:
		p.State[args[2]] = values[0] * values[1]
	case INPUT:
		p.State[args[0]] = p.Input()
	case OUTPUT:
		p.Output(values[0])
	case JUMP_IF_TRUE:
		if values[0] != 0 {
			p.pc = values[1] - 3
		}
	case JUMP_IF_FALSE:
		if values[0] == 0 {
			p.pc = values[1] - 3
		}
	case LESS_THAN:
		if values[0] < values[1] {
			p.State[args[2]] = 1
		} else {
			p.State[args[2]] = 0
		}
	case EQUALS:
		if values[0] == values[1] {
			p.State[args[2]] = 1
		} else {
			p.State[args[2]] = 0
		}
	case ADD_RELATIVE:
		p.rb += values[0]
	}
	p.pc += NUM_ARGS[op] + 1
}
