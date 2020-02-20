package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const VERTICAL = 0
const HORIZONTAL = 1

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type point struct {
	x, y int
}
type line struct {
	a, b      point
	distance  int
	direction int
}

func NewLine(start point, distance int, vector string) (*line, int) {
	l := line{a: start, b: start, distance: distance}
	direction := vector[0]
	magnitude, _ := strconv.Atoi(vector[1:])
	switch direction {
	case 'U':
		l.b.y -= magnitude
		l.direction = VERTICAL
	case 'D':
		l.b.y += magnitude
		l.direction = VERTICAL
	case 'L':
		l.b.x -= magnitude
		l.direction = HORIZONTAL
	case 'R':
		l.b.x += magnitude
		l.direction = HORIZONTAL
	}
	return &l, distance + magnitude
}
func (this line) intersects(other line) (point, bool) {
	if this.direction == VERTICAL && other.direction == HORIZONTAL {
		this, other = other, this
	}

	for _, l := range []*line{&this, &other} {
		if l.a.x > l.b.x || l.a.y > l.b.y {
			l.a, l.b = l.b, l.a
			defer func() { l.a, l.b = l.b, l.a }()
		}
	}

	if this.direction != other.direction &&
		this.a.x <= other.a.x && this.b.x >= other.a.x &&
		this.a.y >= other.a.y && this.a.y <= other.b.y {
		return point{other.a.x, this.a.y}, true
	}
	return point{0, 0}, false
}

type wire struct {
	start, end point
	segments   []line
}

func NewWire(lines string) *wire {
	var w wire
	var l *line
	w.start = point{0, 0}
	w.end = w.start
	d := 0

	vectors := strings.Split(lines, ",")
	w.segments = make([]line, len(vectors))
	for i := range vectors {
		l, d = NewLine(w.end, d, vectors[i])
		w.end = l.b
		w.segments[i] = *l
	}
	return &w
}

type intersection struct {
	a, b line
	p    point
}

func (this wire) intersections(other *wire) chan intersection {
	c := make(chan intersection)

	go func() {
		for _, a := range this.segments {
			for _, b := range other.segments {
				p, found := a.intersects(b)
				if found && (p.x != 0 || p.y != 0) {
					c <- intersection{a, b, p}
				}
			}
		}
		close(c)
	}()

	return c
}

func (this wire) intersection_distances(other *wire) chan int {
	c := make(chan int)

	go func() {
		for i := range this.intersections(other) {
			total_distance := i.a.distance + i.b.distance +
				Abs(i.a.a.x-i.p.x) + Abs(i.a.a.y-i.p.y) +
				Abs(i.b.a.x-i.p.x) + Abs(i.b.a.y-i.p.y)
			c <- total_distance
		}
		close(c)
	}()

	return c
}

func main() {

	file, _ := os.Open("3/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	wire1 := NewWire(scanner.Text())
	//wire1 := NewWire("R8,U5,L5,D3")
	scanner.Scan()
	wire2 := NewWire(scanner.Text())
	//wire2 := NewWire("U7,R6,D4,L4")

	closest := point{0x3fffffffffffffff, 0x3fffffffffffffff}
	for i := range wire1.intersections(wire2) {
		if Abs(i.p.x)+Abs(i.p.y) < Abs(closest.x)+Abs(closest.y) {
			closest = i.p
		}
	}
	fmt.Println(Abs(closest.x) + Abs(closest.y))

	shortest := 0x7fffffffffffffff
	for distance := range wire1.intersection_distances(wire2) {
		if distance < shortest {
			shortest = distance
		}
	}
	fmt.Println(shortest)
}
