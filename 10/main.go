package main

import (
	"../geometry"
	"fmt"
)

const DEBUG = false

type AsteroidField struct {
	geometry.Board
	asteroids [][]bool
}

func newAsteroidField(filename string) *AsteroidField {
	b, err := geometry.NewBoard(filename)
	if err != nil {
		fmt.Println("Error:", err)
	}
	f := AsteroidField{*b, make([][]bool, len(b.Data))}

	// cache string tests
	for x := 0; x < len(b.Data[0]); x++ {
		f.asteroids[x] = make([]bool, len(b.Data))
		for y := 0; y < len(b.Data); y++ {
			f.asteroids[x][y] = b.Data[x][y] == '#'
		}
	}

	return &f
}
func (f AsteroidField) Print() {
	for y := 0; y < f.H; y++ {
		for x := 0; x < f.W; x++ {
			if f.asteroids[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}

func (f AsteroidField) isAsteroid(p geometry.Point) bool {
	return f.asteroids[p.X][p.Y]
}

func (f AsteroidField) vaporizeIfAsteroid(p geometry.Point) bool {
	if f.isAsteroid(p) {
		return true
	}
	return false
}

func (f AsteroidField) anyAsteroidsLeft() bool {
	for _, line := range f.asteroids {
		for _, a := range line {
			if a {
				return true
			}
		}
	}
	return false
}

func bestAsteroid(filename string) (max int, best geometry.Point) {
	f := newAsteroidField(filename)

	max = 0
	for center := range f.Iterate() {
		if !f.isAsteroid(center) {
			continue
		}
		count := 0
		for p := range f.GetRayCollisions(center, f.isAsteroid) {
			_ = p
			count++
		}
		if count > max {
			max = count
			best = center
		}
	}
	return max, best
}

func vaporize(filename string, station geometry.Point) <-chan geometry.Point {
	f := newAsteroidField(filename)
	f.asteroids[station.X][station.Y] = false
	ch := make(chan geometry.Point)

	go func() {
		for f.anyAsteroidsLeft() {
			for p := range f.GetRayCollisions(station, f.isAsteroid) {
				f.asteroids[p.X][p.Y] = false
				ch <- p
			}
		}

		close(ch)
	}()
	return ch
}

func main() {
	//geometry.Box{25, 25}.TestCast(geometry.Point{0, 0}, func(p geometry.Point) bool { return false })
	//fmt.Println()
	//geometry.Box{25, 25}.TestCast(geometry.Point{0, 24}, func(p geometry.Point) bool { return false })
	//fmt.Println()
	//geometry.Box{25, 25}.TestCast(geometry.Point{24, 24}, func(p geometry.Point) bool { return false })
	//fmt.Println()
	//geometry.Box{25, 25}.TestCast(geometry.Point{24, 0}, func(p geometry.Point) bool { return false })
	//fmt.Println()
	//geometry.Box{25, 25}.TestCast(geometry.Point{10, 20}, func(p geometry.Point) bool { return false })
	//fmt.Println()
	//geometry.Box{25, 25}.TestCast(geometry.Point{13, 3}, func(p geometry.Point) bool { return false })
	//fmt.Println()

	// A
	var max int
	var station geometry.Point
	if DEBUG {
		fmt.Println(bestAsteroid("./10/test1.txt"))
		fmt.Println(bestAsteroid("./10/test2.txt"))
		fmt.Println(bestAsteroid("./10/test3.txt"))
		max, station = bestAsteroid("./10/test4.txt")
		fmt.Println(max, station)
	} else {
		max, station = bestAsteroid("./10/input.txt")
		fmt.Println(max)
	}

	// B
	if DEBUG {
		fmt.Println()
		known := map[int]bool{1: true, 2: true, 3: true, 10: true, 20: true, 50: true, 100: true, 199: true, 200: true, 201: true, 299: true}
		i := 1
		for p := range vaporize("./10/test4.txt", station) {
			if known[i] {
				fmt.Println(i, p)
			}
			i++
		}
	} else {
		i := 1
		for p := range vaporize("./10/input.txt", station) {
			if i == 200 {
				fmt.Println(p.X*100 + p.Y)
				break
			}
			i++
		}
	}
}
