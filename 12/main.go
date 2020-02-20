package main

import (
	"../geometry"
	"../utils"
	"bufio"
	"fmt"
	"os"
)

const DEBUG = false

type moon struct {
	geometry.Point
	v geometry.Vector
}
type system struct {
	moons []moon
}

func newMoon(coords string) moon { // awooo
	//coords = "<x=14, y=2, z=8>"
	var m moon
	fmt.Sscanf(coords, "<x=%d, y=%d, z=%d>", &m.X, &m.Y, &m.Z)
	return m
}
func newSystem(filename string) system {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	i := 0
	moons := make([]moon, 4)
	for scanner.Scan() {
		moons[i] = newMoon(scanner.Text())
		i++
	}
	return system{moons}
}

func (a *moon) calculateGravity(b moon) {
	a.v.X += utils.Limit(b.X-a.X, 1)
	a.v.Y += utils.Limit(b.Y-a.Y, 1)
	a.v.Z += utils.Limit(b.Z-a.Z, 1)
}
func (a *moon) totalEnergy() (energy int) {
	potential := a.DistanceToOrigin().Sum()
	kinetic := a.v.Abs().Sum()
	if DEBUG {
		fmt.Println(a.Point, potential, a.v, kinetic)
	}
	return potential * kinetic
}

func (s system) calculateGravity() {
	for i := range s.moons {
		for j := range s.moons {
			if i == j {
				continue
			}
			s.moons[i].calculateGravity(s.moons[j])
		}
	}
}
func (s system) applyVelocity() {
	for i := range s.moons {
		s.moons[i].Add(s.moons[i].v)
	}
}
func (s system) totalEnergy() (sum int) {
	for _, moon := range s.moons {
		sum += moon.totalEnergy()
	}
	return sum
}
func (s system) print() {
	for _, moon := range s.moons {
		fmt.Println(moon.Point, moon.v)
	}
	fmt.Println()
}
func (s system) run(steps int) {
	if DEBUG {
		s.print()
	}
	for i := 1; i <= steps; i++ {
		s.calculateGravity()
		s.applyVelocity()
		if DEBUG {
			fmt.Println(i)
			s.print()
		}
	}
}
func (s system) findPeriods() int {
	// a period is found when both x and v.x repeat
	const X = 0
	const Y = 1
	const Z = 2
	periods := make([]int, 3)
	type pair struct{ p, v *int }

	// Create arrays of pointers to the values, so we don't have to construct the pair every time
	startMoons := make([]moon, len(s.moons))
	copy(startMoons, s.moons)
	start := make([][]pair, 3)
	current := make([][]pair, 3)
	for i := range start {
		start[i] = make([]pair, len(startMoons))
		current[i] = make([]pair, len(startMoons))
	}
	for i := range startMoons {
		start[X][i] = pair{&startMoons[i].X, &startMoons[i].v.X}
		start[Y][i] = pair{&startMoons[i].Y, &startMoons[i].v.Y}
		start[Z][i] = pair{&startMoons[i].Z, &startMoons[i].v.Z}
	}
	for i := range s.moons {
		current[X][i] = pair{&s.moons[i].X, &s.moons[i].v.X}
		current[Y][i] = pair{&s.moons[i].Y, &s.moons[i].v.Y}
		current[Z][i] = pair{&s.moons[i].Z, &s.moons[i].v.Z}
	}
	samePosition := func(axis int) bool {
		for i := range start {
			if *(start[axis][i].p) != *(current[axis][i].p) ||
				*(start[axis][i].v) != *(current[axis][i].v) {
				return false
			}
		}
		return true
	}

	date := 1
	for periods[X] == 0 || periods[Y] == 0 || periods[Z] == 0 {
		s.calculateGravity()
		s.applyVelocity()
		for axis, p := range periods {
			if p != 0 {
				continue
			}
			if samePosition(axis) {
				periods[axis] = date
			}
		}

		date++
	}

	return utils.LCM(periods[X], utils.LCM(periods[Y], periods[Z]))
}

func main() {
	// PART 1
	if DEBUG {
		system := newSystem("12/test1.txt")
		system.run(10)
		fmt.Println(system.totalEnergy())

		system = newSystem("12/test2.txt")
		system.run(100)
		fmt.Println(system.totalEnergy())
	} else {
		system := newSystem("12/input.txt")
		system.run(1000)
		fmt.Println(system.totalEnergy())
	}

	// PART 2
	if DEBUG {
		system := newSystem("12/test1.txt")
		fmt.Println(system.findPeriods())
		system = newSystem("12/test2.txt")
		fmt.Println(system.findPeriods())
	} else {
		system := newSystem("12/input.txt")
		fmt.Println(system.findPeriods())
	}
}
